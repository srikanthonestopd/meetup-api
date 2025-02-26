package routes

import (
	"github.com/gorilla/mux"
	"meetup-apis/handlers"
)

func RegisterItemRoutes(router *mux.Router) {
	router.HandleFunc("/api/payment", handlers.PaymentsData).Methods("POST")
	router.HandleFunc("/api/addToCart", handlers.AddToCartData).Methods("POST")
	router.HandleFunc("/api/get/{id}", handlers.GetData).Methods("GET")
	router.HandleFunc("/api/next-id", handlers.GetNextItemIDHandler).Methods("GET")

}
