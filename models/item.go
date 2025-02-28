package models

import "time"

// Item struct for Couchbase storage
type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Payment struct {
	Username       string  `json:"username"`
	UserEmailID    string  `json:"user_emailid"`
	EventName      string  `json:"event_name"`
	EventId        string  `json:"event_id"` // Ensure correct casing
	Price          float64 `json:"price"`
	Quantity       int     `json:"quantity"`
	TotalPrice     float64 `json:"total_price"`
	FullName       string  `json:"full_name"`
	BillingAddress string  `json:"billing_address"`
	PhoneNumber    string  `json:"phone_number"`
	CardNumber     string   `json:"card_number"`
	NameOnCard     string  `json:"name_on_card"`
	CardType       string  `json:"card_type"`
	ExpiryMM       string  `json:"expiry_mm"`
	ExpiryYYYY     string  `json:"expiry_yyyy"`
	CVV            string  `json:"cvv"`
}

type CartData struct {
	UserName       string    `json:"username"`
	UserEmailID    string    `json:"email_id"`
	EventName      string    `json:"event_name"`
	EventId        string    `json:"event_id"`
	Price          float64   `json:"price"`
	Quantity       int       `json:"quantity"`
	TotalPrice     float64   `json:"total_price"`
	FullName       string    `json:"full_name"`
	PhoneNumber    string    `json:"phone_number"`
	BillingAddress string    `json:"billing_address"`
	CardNumber     string       `json:"card_number"`
	NameOnCard     string    `json:"name_on_card"`
	CardType       string    `json:"card_type"`
	Expiry_mm      string    `json:"expiry_mm"`
	Expiry_yyyy    string    `json:"expiry_yyyy"`
	CVV            string    `json:"cvv"`
	PurchasedOn    time.Time `json:"purchased_on"`
}
type RegisterData struct{
    FullName        string      `json:"full_name"`
    EmailID         string      `json:"email_id"`
    Password        string      `json:"password"`
}

type CreateProfile struct{
    FullName        string      `json:"full_name"`
    EmailID         string      `json:"email_id"`
    Password        string      `json:"password"`
    CreatedOn       time.Time   `json:"created_on"`
}

type BarcodeData struct {
	EventId   string `json:"event_id"`
	Barcode   string `json:"barcode"`
	ScannerID string `json:"scanner_id"`
	ScanTime  string `json:"scan_time"`
	Location  string `json:"location"`
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
}
type NotificationsData struct{
    UserID              string          `json:"user_id"`
    UserName            string          `json:"user_name"`
    MessageType         string          `json:"message_type"`
    EmailId             string          `json:"email_id"`
    Subject             string          `json:"subject"`
    Message             string          `json:"message"`
}
type InviteData struct{
    UserName            string           `json:"User_name"`
    EventId             string           `json:"event_id"`
    EventName           string           `json:"event_name"`
    Message             string           `json:"message"`
    MessageType         string           `json:"message_type"`
    EmailId             string           `json:"email_id"`
}