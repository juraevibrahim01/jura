package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/service"
)

type Auth_handler struct {
	service *service.Auth_service
}

func Auth_new_handler(service *service.Auth_service) *Auth_handler {
	return &Auth_handler{service: service}
}

// ---------------------------------------------- login -----------------------------------------------------------------------------------
func (h *Auth_handler) Login(w http.ResponseWriter, r *http.Request) {

	// Response | Request
	var req models.Auth_Req_Login
	var res models.Auth_Res_Login

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	// Проверка: Все поля возвращются
	if req.Email == "" || req.Password == "" {
		res.Status = "error"
		res.Error = "Поля не валидны"
		w.WriteHeader(400)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	// Проверка на идентификацию
	id_user, err := h.identification(&req.Email)
	if err != nil || id_user == 0 {
		res.Status = "error"
		res.Error = "Ошибка при вводе логина и пароля"
		w.WriteHeader(404)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	// Проверка пароля
	err = h.check_password(&id_user, &req.Password)
	if err != nil {
		res.Status = "error"
		res.Error = "Ошибка при вводе логина и пароля"
		w.WriteHeader(404)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	// otp
	err = h.otp(&req.Email)
	if err != nil {
		res.Status = "error"
		res.Description = "Ошибка отправки OTP на email"
		w.WriteHeader(500)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Status = "success"
	res.Step = "otp"
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

func (h *Auth_handler) identification(email *string) (int, error) {
	id_user, err := h.service.Service_identification(email)
	if err != nil {
		return 0, err
	}
	return id_user, nil
}

func (h *Auth_handler) check_password(id *int, password *string) error {
	err := h.service.Service_check_password(id, password)
	if err != nil {
		return err
	}
	return nil
}

func (h *Auth_handler) otp(email *string) error {
	otpCode := h.service.GenerateOTP()
	h.service.SaveOTP(email, &otpCode)

	// Отпаврка в почту
	err := h.service.SendOTPEmail(*email, otpCode)
	if err != nil {
		return err
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------------------

// ---------------------------------------------- otp -----------------------------------------------------------
func (h *Auth_handler) Check_otp(w http.ResponseWriter, r *http.Request) {

	// Response | Request
	var req models.Auth_req_otp
	var res models.Auth_res_otp

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	verefi := h.service.OtpVerify(req.Email, req.Code)
	if verefi != true {
		res.Status = "error"
		res.Error = "Отп просрочен или не верный"
		w.WriteHeader(400)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	userID, err := h.service.Service_identification(&req.Email)
	if err != nil {
		res.Status = "error"
		res.Error = "Не удалось определить пользователя"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(res)
		return
	}

	accessToken, refreshToken, err := h.service.GenerationToken(userID, &req.Email, req.Role)
	if err != nil {
		log.Print("Ошибка при генерации токенов: ", err)
		res.Status = "error"
		res.Error = "Не удалось сгенерировать токены"
		w.WriteHeader(500)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Status = "success"
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	w.WriteHeader(200)
	// response
	json.NewEncoder(w).Encode(res)
}

// ------------------------------------------------------------------

// --------------------------- ref_token -----------------------------
func (h *Auth_handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	// Response | Request
	var req models.Auth_req_ref_token
	var res models.Auth_res_ref_token

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	claims, err := h.service.ValidateToken("", req.RefToken)
	if err != nil {
		log.Println("Ошибка валидации токена:", err)
		w.WriteHeader(http.StatusUnauthorized)
		res.Error = "Недействительный токен"
		res.Status = "error"
		json.NewEncoder(w).Encode(res)
		return
	}

	token, ref_token, err := h.service.GenerationToken(claims.UserID, &claims.Email, claims.Role)
	if err != nil {
		log.Print("Ошибка при генерации токенов")
		w.WriteHeader(500)
		res.Status = "error"
		res.Error = "Ошибка сервера"
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(200)
	res.Status = "success"
	res.Token = token
	res.RefToken = ref_token
	json.NewEncoder(w).Encode(res)

}

// -----------------------------------------------------------------
