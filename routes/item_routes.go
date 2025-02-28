package routes

import (
	"github.com/gorilla/mux"
	"meetup-apis/handlers"
)

func RegisterItemRoutes(router *mux.Router) {
	router.HandleFunc("/api/payment/checkout", handlers.PaymentsData).Methods("POST")
	router.HandleFunc("/api/auth/register", handlers.RegisterData).Methods("POST")
	router.HandleFunc("/api/tickets/verify/{barcode}", handlers.BarcodeData).Methods("POST")
	router.HandleFunc("/api/auth/createProfile", handlers.CreateProfile).Methods("POST")
	router.HandleFunc("/api/cart/add", handlers.AddtocartData).Methods("POST")
	router.HandleFunc("/api/get/{id}", handlers.GetData).Methods("GET")
	router.HandleFunc("/api/next-id", handlers.GetNextItemIDHandler).Methods("GET")

}
