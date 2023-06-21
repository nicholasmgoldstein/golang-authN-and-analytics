package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/myusername/auth-service/controllers"
	"github.com/myusername/auth-service/middleware"
)

func SetAuthRoutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST").Queries("all", "{all}")
	router.HandleFunc("/refresh_token", controllers.RefreshToken).Methods("POST")
	router.HandleFunc("/verify_email", controllers.VerifyEmail).Methods("GET")
	router.HandleFunc("/reset_password", controllers.ResetPasswordRequest).Methods("POST")
	router.HandleFunc("/reset_password/{token}", controllers.ResetPassword).Methods("POST")

	router.Use(middleware.LogRequest)
}
