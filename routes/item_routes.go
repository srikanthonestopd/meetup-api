package routes

import (
	"meetup-apis/handlers"

	"github.com/gorilla/mux"
)

func RegisterItemRoutes(router *mux.Router) {
	router.HandleFunc("/api/payment", handlers.PaymentsData).Methods("POST")
	//router.HandleFunc("/api/auth/login", handlers.LoginData).Methods("POST")
	router.HandleFunc("/api/auth/register", handlers.RegisterData).Methods("POST")
	router.HandleFunc("/api/auth/createProfile", handlers.CreateProfile).Methods("POST")
	router.HandleFunc("/api/auth/forgotpassword", handlers.ForgotPassword).Methods("POST")
	router.HandleFunc("/api/events/create", handlers.CreateEvent).Methods("POST")
	router.HandleFunc("/api/addToCart", handlers.AddToCartData).Methods("POST")
	router.HandleFunc("/api/Cart/Checkout", handlers.Checkout).Methods("POST")
	router.HandleFunc("/api/payment/refund/{id}", handlers.RefundData).Methods("POST")
	router.HandleFunc("/api/tickets/cancel/{id}", handlers.TicketCancellation).Methods("POST")
	router.HandleFunc("/api/share/event/{id}", handlers.ShareEvent).Methods("POST")
	router.HandleFunc("/api/ads/click/{id}", handlers.TrackClickData).Methods("POST")
	router.HandleFunc("/api/reviews/add", handlers.ReviewsData).Methods("POST")
	router.HandleFunc("/api/get/{id}", handlers.GetData).Methods("GET")
	router.HandleFunc("/api/next-id", handlers.GetNextItemIDHandler).Methods("GET")

}
