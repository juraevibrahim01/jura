package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/smtp"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type Auth_service struct {
	repository *repository.Auth_reposirory
}

func Auth_new_service(repository *repository.Auth_reposirory) *Auth_service {
	return &Auth_service{repository: repository}
}

// ---------------------------- login ------------------------------------
// Проверка на идентификации
func (s *Auth_service) Service_identification(email *string) (int, error) {
	id_user, err := s.repository.Reposirory_identification(email)
	if err != nil {
		log.Print("Ошибка: ", err)
		return 0, err
	}
	return id_user, nil
}

// Проверка на совподение пороля
func (s *Auth_service) Service_check_password(id *int, password *string) error {
	hash_password, err := s.repository.Reposirory_check_password(id)
	if err != nil {
		return err
	}
	err = s.check_password_hash(&hash_password, password)
	if err != nil {
		return err
	}
	return nil
}

// Проверка на соответствия хеша и пороля из фронт
func (s *Auth_service) check_password_hash(hash_password, password *string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*hash_password), []byte(*password))
	if err != nil {
		log.Print("Ошибка при дехешировании: ", err)
		return err
	}
	return nil
}

// ----------------------------------------------------------------------

// -------------------------------------- otp -------------------------------

// Защищаем карту с помощью мьютекса
var (
	otpCode  = make(map[string]string)
	otpMutex sync.RWMutex
)

// Безопасная генерация 6-значного OTP через crypto/rand
func (s *Auth_service) GenerateOTP() string {
	max := big.NewInt(1000000)
	nBig, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "000000" // Деградация до дефолта при ошибке энтропии
	}
	return fmt.Sprintf("%06d", nBig.Int64())
}

// Сохранение с потокобезопасным удалением
func (s *Auth_service) SaveOTP(email, otp *string) {
	otpMutex.Lock()
	otpCode[*email] = *otp
	otpMutex.Unlock()

	go func() {
		time.Sleep(5 * time.Minute)
		otpMutex.Lock()
		delete(otpCode, *email)
		otpMutex.Unlock()
	}()
}

// Проверка пароля с удалением после успешной проверки (Одноразовый код)
func (s *Auth_service) OtpVerify(email, otp string) bool {
	otpMutex.Lock() // Используем Lock, так как будем удалять код при успехе
	defer otpMutex.Unlock()

	savedCode, exists := otpCode[email]
	log.Println("saved code: ", savedCode, "provided otp:  ", otp)

	if !exists {
		return false
	}

	if savedCode == otp {
		delete(otpCode, email) // Удаляем код, чтобы его нельзя было использовать дважды
		return true
	}

	return false
}

func (s *Auth_service) SendOTPEmail(toEmail string, otp string) error {
	from := "pharmpro111@gmail.com"
	password := "cqqavqhjabksjdfd"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Ваш OTP код"
	body := fmt.Sprintf("Ваш код: %s\nОн действителен 5 минут.", otp)
	msg := []byte("From: " + from + "\r\n" + "To: " + toEmail + "\r\n" + "Subject: " + subject + "\r\n" + "\r\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, msg)
	if err != nil {
		log.Println("SMTP ошибка:", err)
		return err
	}

	fmt.Println("Письмо успешно отправлено на", toEmail)
	return nil
}

func (s *Auth_service) GenerationToken(email *string) (string, string, error) {
	claimsToken := jwt.MapClaims{
		"email": *email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}

	claimsRefToken := jwt.MapClaims{
		"email": *email,
		"exp":   time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)
	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefToken)

	// находим ключь определенного юзера
	access_token, ref_token, err := s.repository.Repository_choose_otpkey(email)
	if err == nil {
		log.Println("Не удалось получать токены из бд!")
		return "", "", errors.New("Не удалось получить токены из бд!")
	}

	// Предполагается, что модели хранят []byte секреты
	accessToken, err := token.SignedString(access_token)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := refToken.SignedString(ref_token)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *Auth_service) ValidateToken(tokenString, refTokenString, email string) (*models.Claims, error) {
	var originToken string
	var secret []byte

	if tokenString != "" {
		originToken = strings.TrimPrefix(tokenString, "Bearer ")
		access_token, _, err := s.repository.Repository_choose_otpkey(&email)
		if err != nil {
			log.Println("Не удалось получать access-токен из бд!", err)
			return nil, errors.New("Не удалось получить токены из бд!")
		}
		secret = []byte(access_token)
	} else if refTokenString != "" {
		originToken = refTokenString
		_, ref_token, err := s.repository.Repository_choose_otpkey(&email)
		if err != nil {
			log.Println("Не удалось получать refresh-токен из бд!", err)
			return nil, errors.New("Не удалось получить токены из бд!")
		}
		secret = []byte(ref_token)
	} else {
		return nil, errors.New("токен не передан")
	}

	token, err := jwt.ParseWithClaims(originToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи")
		}
		return secret, nil
	})

	if err != nil {
		log.Print("недействительный токен")
		return nil, models.ErrTokenInvalid
	}

	if !token.Valid {
		log.Print("недействительный токен")
		return nil, models.ErrTokenInvalid
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		log.Print("не удалось получить данные из токена")
		return nil, errors.New("не удалось получить данные из токена")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		log.Print("токен истек")
		return nil, models.ErrTokenExpired
	}

	return claims, nil
}

// ---------------------------------------------------------------
