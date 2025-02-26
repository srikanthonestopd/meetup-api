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
