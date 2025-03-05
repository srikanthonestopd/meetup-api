package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"meetup-apis/config"
	"meetup-apis/models"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

// Fetch all items, find the highest ID, and increment it
func getNextItemID() (string, error) {
	query := "SELECT id FROM `roh-api`.`myscope`.`mycollection`;"
	rows, err := config.Cluster.Query(query, nil)
	if err != nil {
		return "", err
	}

	var ids []int
	for rows.Next() {
		var result map[string]interface{}
		err := rows.Row(&result)
		if err != nil {
			return "", err
		}

		// Convert ID to int (assuming numeric IDs)
		if idStr, ok := result["id"].(string); ok {
			var id int
			_, err := fmt.Sscanf(idStr, "item%d", &id)
			if err == nil {
				ids = append(ids, id)
			}
		}
	}

	// Sort and find the last ID
	sort.Ints(ids)
	newID := 1
	if len(ids) > 0 {
		newID = ids[len(ids)-1] + 1 // Increment last ID
	}

	return fmt.Sprintf("item%d", newID), nil
}

// Get the next available item ID
func GetNextItemIDHandler(w http.ResponseWriter, r *http.Request) {
	nextID, err := getNextItemID()
	if err != nil {
		http.Error(w, "Failed to get next item ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"next_id": nextID})
}

// Get data from Couchbase by ID
func GetData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Fetch document from Couchbase
	result, err := config.Collection.Get(id, nil)
	if err != nil {
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}

	var item models.Item
	err = result.Content(&item)
	if err != nil {
		http.Error(w, "Failed to parse data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// Insert data into Couchbase with auto-incrementing ID
func PaymentsData(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}

	// Print inserting data
	fmt.Printf("🔹 Inserting payment with Event ID: %s\n", payment.EventId)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(payment.EventId, payment, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data inserted successfully",
		"id":      payment.EventId,
	})
}

func AddToCartData(w http.ResponseWriter, r *http.Request) {
	var cartData models.CartData
	err := json.NewDecoder(r.Body).Decode(&cartData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with ID: %s\n", cartData.EventId)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(cartData.PhoneNumber, cartData, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data inserted successfully",
		"id":      cartData.PhoneNumber,
	})
}

/*func LoginData(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Checking credentials for Email: %s\n", login.Email)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Retrieving user data from Couchbase
	getResult, err := meetupCollection.Get(login.Email, nil)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		log.Println("❌ User not found:", err)
		return
	}

	var register models.RegisterData
	err = getResult.Content(&register)
	if err != nil {
		http.Error(w, "Failed to retrieve user details", http.StatusInternalServerError)
		log.Println("❌ Failed to retrieve user details:", err)
		return
	}

	if register.Password != login.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		log.Println("❌ Invalid Password or Email:", login.Email)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login Successful",
		"id":      login.Email,
	})
}*/

func RegisterData(w http.ResponseWriter, r *http.Request) {
	var registerData models.RegisterData
	err := json.NewDecoder(r.Body).Decode(&registerData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}

	// Print inserting data
	fmt.Printf("🔹 Inserting item with EmailID: %s\n", registerData.EmailID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(registerData.EmailID, registerData, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data inserted successfully",
		"id":      registerData.EmailID,
	})
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var forgotPassword models.ForgotPassword
	fmt.Println("🔹 Inserting item with EmailID:")

	err := json.NewDecoder(r.Body).Decode(&forgotPassword)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with EmailID: %s\n", forgotPassword.Email)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(forgotPassword.Email, forgotPassword, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "New Password Created successfully",
		"id":      forgotPassword.Email,
	})
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	var createProfile models.CreateProfile
	fmt.Println("🔹 Inserting item with EmailID:")

	err := json.NewDecoder(r.Body).Decode(&createProfile)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with EmailID: %s\n", createProfile.EmailID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(createProfile.EmailID, createProfile, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data inserted successfully",
		"id":      createProfile.EmailID,
	})
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var createEvent models.CreateEvent
	fmt.Println("🔹 Inserting item with EventID:")

	err := json.NewDecoder(r.Body).Decode(&createEvent)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with EmailID: %s\n", createEvent.ID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(createEvent.ID, createEvent, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Event Created successfully",
		"id":      createEvent.ID,
	})
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	var checkout models.Checkout
	fmt.Println("🔹 Inserting item with EventID:")

	err := json.NewDecoder(r.Body).Decode(&checkout)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with EventID: %s\n", checkout.EventID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(checkout.EventID, checkout, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Checkout successful. Proceed to payment to complete your order.",
		"id":      checkout.EventID,
	})
}

func RefundData(w http.ResponseWriter, r *http.Request) {
	var refundData models.RefundData
	fmt.Println("🔹 Inserting item with UserID:")

	err := json.NewDecoder(r.Body).Decode(&refundData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with UserID: %s\n", refundData.UserID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(refundData.UserID, refundData, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Refund Successful",
		"id":      refundData.UserID,
	})
}

func TicketCancellation(w http.ResponseWriter, r *http.Request) {
	var ticketCancellation models.CancelTicketRequest
	fmt.Println("🔹 Inserting item with UserID:")

	err := json.NewDecoder(r.Body).Decode(&ticketCancellation)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with UserID: %s\n", ticketCancellation.UserID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(ticketCancellation.UserID, ticketCancellation, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Ticket Booking Cancelled",
		"id":      ticketCancellation.UserID,
	})
}

func ReviewsData(w http.ResponseWriter, r *http.Request) {
	var reviewData models.ReviewData
	fmt.Println("🔹 Inserting item with UserID:")

	err := json.NewDecoder(r.Body).Decode(&reviewData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with UserID: %s\n", reviewData.UserID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(reviewData.UserID, reviewData, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Thank you for your valuable feedback",
		"id":      reviewData.UserID,
	})
}

func ShareEvent(w http.ResponseWriter, r *http.Request) {
	var shareEvent models.ShareEvent
	fmt.Println("🔹 Inserting item with UserID:")

	err := json.NewDecoder(r.Body).Decode(&shareEvent)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with UserID: %s\n", shareEvent.UserID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(shareEvent.UserID, shareEvent, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Event Shared Successfully",
		"id":      shareEvent.UserID,
	})
}

func TrackClickData(w http.ResponseWriter, r *http.Request) {
	var clickData models.TrackClickData
	fmt.Println("🔹 Inserting item with UserID:")

	err := json.NewDecoder(r.Body).Decode(&clickData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}
	// Print inserting data
	fmt.Printf("🔹 Inserting item with UserID: %s\n", clickData.UserID)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(clickData.UserID, clickData, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data Inserted Successfully",
		"id":      clickData.UserID,
	})
}
