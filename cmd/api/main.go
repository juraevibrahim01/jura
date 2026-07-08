package main

import (
	"log"
	"net/http"

	"github.com/juraevibrahim01/jura/internal/handler"
	"github.com/juraevibrahim01/jura/internal/middleware"
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

	// ---------------------------------- apis --------------------------------------
	// Маршрутизатор
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", auth_handler.Login)
	mux.HandleFunc("POST /login/check_otp", auth_handler.Check_otp)
	mux.Handle("POST /user", middleware.AuthMiddleware(auth_service, http.HandlerFunc(user_handler.User_create)))
	mux.Handle("GET /tickets", middleware.AuthMiddleware(auth_service, middleware.RoleMiddleware(user_service, []string{"reading", "admin"}, http.HandlerFunc(ticket_handler.GetTickets))))

	handleWithCors := middleware.CORSMiddleware(mux)

	// url
	log.Fatal(http.ListenAndServe(":8081", handleWithCors))
	// log.Fatal - если порт занят то программа не промолчит а даст информацию что порт занят
}
