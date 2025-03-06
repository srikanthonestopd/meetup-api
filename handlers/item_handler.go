package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	"meetup-apis/config"
	"meetup-apis/models"
	"net/http"
	"sort"
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

// Get data from Couchbase by ID
func GetAddtocartData(w http.ResponseWriter, r *http.Request) {
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

func AddtocartData(w http.ResponseWriter, r *http.Request) {
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

// Insert data into Couchbase with auto-incrementing ID
func BarcodeData(w http.ResponseWriter, r *http.Request) {
	var barcode models.BarcodeData
	err := json.NewDecoder(r.Body).Decode(&barcode)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}

	// Print inserting data
	fmt.Printf("🔹 Inserting barcode with Event ID: %s\n", barcode.EventId)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(barcode.EventId, barcode, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data inserted successfully",
		"id":      barcode.EventId,
	})
}

// Insert data into Couchbase with auto-incrementing ID
func NotificationsData(w http.ResponseWriter, r *http.Request) {
	var notifications models.NotificationsData
	err := json.NewDecoder(r.Body).Decode(&notifications)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}

	// Print inserting data
	fmt.Printf("🔹 Inserting notifications with Email ID: %s\n", notifications.EmailId)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(notifications.EmailId, notifications, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data inserted successfully",
		"id":      notifications.EmailId,
	})
}

// Insert data into Couchbase with auto-incrementing ID
func InviteData(w http.ResponseWriter, r *http.Request) {
	var invite models.InviteData
	err := json.NewDecoder(r.Body).Decode(&invite)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("❌ JSON Decode Error:", err)
		return
	}

	// Print inserting data
	fmt.Printf("🔹 Inserting invite data with Email ID: %s\n", invite.EmailId)
	meetupCollection := config.Cluster.Bucket("roh-api").Scope("myscope").Collection("meetup")

	// Insert into Couchbase
	_, err = meetupCollection.Insert(invite.EmailId, invite, nil)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		log.Println("❌ Couchbase Insert Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data inserted successfully",

    })
}
