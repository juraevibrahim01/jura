package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/juraevibrahim01/jura/internal/handler"
	"github.com/juraevibrahim01/jura/internal/middleware"
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/repository"
	"github.com/juraevibrahim01/jura/internal/service"
	"github.com/juraevibrahim01/jura/pkg"
)

func main() {
	db, err := pkg.InitPostgres()
	if err != nil {
		log.Print("Ошибка сервера при соединении бд: ", err)
		return
	}
	defer db.DB.Close()

	// ---------------------------------- auth ------------------------------------
	auth_repository := repository.Auth_new_repository(db)
	auth_service := service.Auth_new_service(auth_repository)
	auth_handler := handler.Auth_new_handler(auth_service)

	// ---------------------------------- user ------------------------------------
	user_repository := repository.User_new_repository(db)
	user_service := service.User_new_service(user_repository)
	user_handler := handler.User_new_handler(user_service)

	// ---------------------------------- ticket -----------------------------------
	ticket_repository := repository.Ticket_new_repository(db)
	ticket_service := service.Ticket_new_service(ticket_repository)
	ticket_handler := handler.Ticket_new_handler(ticket_service)

	// ---------------------------------- test_keys -----------------------------------
	test_keys_repository := repository.New_Test_keys_repository(db)
	test_keys_service := service.New_Test_keys_service(test_keys_repository)
	test_keys_handler := handler.New_Test_keys_handler(test_keys_service)

	// ---------------------------------- ai -----------------------------------

	// Загружаем .env до того, как обращаемся к os.Getenv
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Замечание: файл .env не найден, считываем системные переменные")
	}

	aiConfig := models.GeminiConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
		APIURL: "https://generativelanguage.googleapis.com", // Базовый домен
	}

	ai_service := service.NewAI_service(aiConfig)
	ai_handler := handler.NewAIHandler(ai_service)

	// Проверка при старте
	if err := ai_service.HealthCheck(); err != nil {
		log.Printf("AI service warning: %v", err)
	} else {
		log.Println("AI service is healthy")
	}

	// ---------------------------------- apis --------------------------------------
	// Маршрутизатор (gorilla/mux для поддержки методов и параметров)
	router := mux.NewRouter()

	router.HandleFunc("/login", auth_handler.Login).Methods("POST")
	router.HandleFunc("/login/check_otp", auth_handler.Check_otp).Methods("POST")
	router.Handle("/user", middleware.AuthMiddleware(auth_service, http.HandlerFunc(user_handler.User_create))).Methods("POST")
	router.Handle("/tickets", middleware.AuthMiddleware(auth_service, middleware.RoleMiddleware(user_service, []string{"reading", "admin"}, http.HandlerFunc(ticket_handler.GetTickets)))).Methods("GET")
	router.Handle("/tickets", middleware.AuthMiddleware(auth_service, middleware.RoleMiddleware(user_service, []string{"writing", "admin"}, http.HandlerFunc(ticket_handler.Ticket_create)))).Methods("POST")
	router.Handle("/test-keys", middleware.AuthMiddleware(auth_service, middleware.RoleMiddleware(user_service, []string{"reading", "admin"}, http.HandlerFunc(test_keys_handler.GetTestKeys)))).Methods("GET")
	router.Handle("/test-keys", middleware.AuthMiddleware(auth_service, middleware.RoleMiddleware(user_service, []string{"writing", "admin"}, http.HandlerFunc(test_keys_handler.CreateTestKey)))).Methods("POST")
	router.Handle("/test-keys/{id}", middleware.AuthMiddleware(auth_service, middleware.RoleMiddleware(user_service, []string{"reading", "admin"}, http.HandlerFunc(test_keys_handler.GetTestKeyByID)))).Methods("GET")
	router.Handle("/ai", middleware.AuthMiddleware(auth_service, middleware.RoleMiddleware(user_service, []string{"writing", "admin"}, http.HandlerFunc(ai_handler.Chat)))).Methods("POST")
	handleWithCors := middleware.CORSMiddleware(router)

	// -------------------------------- url --------------------------------------
	// Получение порта из переменной окружения (с дефолтом)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, handleWithCors))
	// log.Fatal - если порт занят то программа не промолчит а даст информацию что порт занят
}
