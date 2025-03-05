package models

import "time"

// Item struct for Couchbase storage
type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ForgotPassword struct {
	UserName        string `json:"user_name"`
	Email           string `json:"email"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	CardNumber     int     `json:"card_number"`
	NameOnCard     string  `json:"name_on_card"`
	CardType       string  `json:"card_type"`
	ExpiryMM       string  `json:"expiry_mm"`
	ExpiryYYYY     string  `json:"expiry_yyyy"`
	CVV            string  `json:"cvv"`
}

type CartData struct {
	UserName       string    `json:"user_name"`
	UserEmailID    string    `json:"user_email_id"`
	EventName      string    `jsonz:"event_name"`
	EventId        string    `json:"event_id"`
	Price          float64   `json:"price"`
	Quantity       int       `json:"quantity"`
	TotalPrice     float64   `json:"total_price"`
	FullName       string    `json:"full_name"`
	PhoneNumber    string    `json:"phone_number"`
	BillingAddress string    `json:"billing_address"`
	CardNumber     int       `json:"card_number"`
	NameOnCard     string    `json:"name_on_card"`
	CardType       string    `json:"card_type"`
	Expiry_mm      string    `json:"expiry_mm"`
	Expiry_yyyy    string    `json:"expiry_yyyy"`
	CVV            string    `json:"cvv"`
	PurchasedOn    time.Time `json:"purchased_on"`
}

type RegisterData struct {
	FullName string `json:"full_name"`
	EmailID  string `json:"email_id"`
	Password string `json:"password"`
}

type CreateProfile struct {
	FullName  string    `json:"full_name"`
	EmailID   string    `json:"email_id"`
	Password  string    `json:"password"`
	CreatedOn time.Time `json:"created_on"`
}

type CreateEvent struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Organizer      string    `json:"organizer"`
	Phone          string    `json:"phone"`
	Location       string    `json:"location"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Country        string    `json:"country"`
	ZipCode        string    `json:"zip_code"`
	Date           string    `json:"date"`
	Time           string    `json:"time"`
	Capacity       int       `json:"capacity"`
	AvailableSeats int       `json:"available_seats"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Checkout struct {
	EventID     string  `json:"event_id"`
	EventName   string  `json:"event_name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
	Tax         float64 `json:"tax"`
	ServiceFee  float64 `json:"service_fee"`
	TotalAmount float64 `json:"total_amount"`
}

type RefundData struct {
	EventID       string    `json:"event_id"`
	EventName     string    `json:"event_name"`
	UserID        string    `json:"user_id"`
	PaymentID     string    `json:"payment_id"`
	RefundAmount  float64   `json:"refund_amount"`
	PaymentMethod string    `json:"payment_method"`
	RequestTime   time.Time `json:"request_time"`
	RefundReason  string    `json:"refund_reason"`
	RefundStatus  string    `json:"refund_status"`
}

type CancelTicketRequest struct {
	EventID            string    `json:"event_id"`
	EventName          string    `json:"event_name"`
	UserID             string    `json:"user_id"`
	PaymentID          string    `json:"payment_id"`
	TicketID           string    `json:"ticket_id"`
	CancellationFee    float64   `json:"cancellation_fee"`
	RefundAmount       float64   `json:"refund_amount"`
	PaymentMethod      string    `json:"payment_method"`
	RequestTime        time.Time `json:"request_time"`
	CancellationReason string    `json:"cancellation_reason"`
	RefundEligible     bool      `json:"refund_eligible"`
	ProcessedBy        string    `json:"processed_by"`
}

type ReviewData struct {
	EventID    string    `json:"event_id"`
	EventName  string    `json:"event_name"`
	UserID     string    `json:"user_id"`
	ReviewID   string    `json:"review_id"`
	Rating     float64   `json:"rating"`
	ReviewText string    `json:"review_text"`
	ReviewTime time.Time `json:"review_time"`
}

type ShareEvent struct {
	UserID     string    `json:"user_id"`
	EventID    string    `json:"event_id"`
	EventName  string    `json:"event_name"`
	Platform   string    `json:"platform"`
	Message    string    `json:"message"`
	SharedTime time.Time `json:"shared_time"`
}

type TrackClickData struct {
	UserID      string  `json:"user_id"`
	EventID     string  `json:"event_id"`
	EventName   string  `json:"event_name"`
	ClickTime   string  `json:"click_time"`
	TotalClicks float64 `json:"total_clicks"`
}
