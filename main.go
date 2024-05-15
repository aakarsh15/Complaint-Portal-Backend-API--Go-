package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// User represents a user in the system
type User struct {
	ID         string      `json:"id"`
	SecretCode string      `json:"secret_code"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Complaints []Complaint `json:"complaints"`
}

// Complaint represents a complaint submitted by a user
type Complaint struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Severity int    `json:"severity"`
}

// DB simulates a database for users and complaints
type DB struct {
	users      map[string]User
	complaints map[string]Complaint
	mutex      sync.RWMutex
}

// NewDB creates a new instance of the database
func NewDB() *DB {
	return &DB{
		users:      make(map[string]User),
		complaints: make(map[string]Complaint),
	}
}

var database = NewDB()

func main() {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/submitComplaint", handleSubmitComplaint)
	http.HandleFunc("/getAllComplaintsForUser", handleGetAllComplaintsForUser)
	http.HandleFunc("/getAllComplaintsForAdmin", handleGetAllComplaintsForAdmin)
	http.HandleFunc("/viewComplaint", handleViewComplaint)
	http.HandleFunc("/resolveComplaint", handleResolveComplaint)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Extract secret code from request
	secretCode := r.URL.Query().Get("secret_code")

	// Retrieve user from database
	user, ok := database.users[secretCode]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return user details
	json.NewEncoder(w).Encode(user)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	// Decode request body into a new user
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate unique ID and secret code
	newUser.ID = generateID()
	newUser.SecretCode = generateSecretCode()

	// Save user to database
	database.mutex.Lock()
	defer database.mutex.Unlock()
	database.users[newUser.SecretCode] = newUser

	// Return newly created user details
	json.NewEncoder(w).Encode(newUser)
}

func handleGetAllComplaintsForUser(w http.ResponseWriter, r *http.Request) {
	// Extract secret code from request
	secretCode := r.URL.Query().Get("secret_code")

	// Retrieve user from database
	user, ok := database.users[secretCode]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return user's complaints
	json.NewEncoder(w).Encode(user.Complaints)
}

func handleGetAllComplaintsForAdmin(w http.ResponseWriter, r *http.Request) {
	// Simulate admin access (no authentication for simplicity)
	var complaints []Complaint
	for _, complaint := range database.complaints {
		complaints = append(complaints, complaint)
	}
	json.NewEncoder(w).Encode(complaints)
}

func handleViewComplaint(w http.ResponseWriter, r *http.Request) {
	// Extract complaint ID from request
	complaintID := r.URL.Query().Get("complaint_id")

	// Retrieve complaint from database
	complaint, ok := database.complaints[complaintID]
	if !ok {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	// Return complaint details
	json.NewEncoder(w).Encode(complaint)
}

func handleResolveComplaint(w http.ResponseWriter, r *http.Request) {
	// Extract complaint ID from request
	complaintID := r.URL.Query().Get("complaint_id")

	// Retrieve complaint from database
	complaint, ok := database.complaints[complaintID]
	if !ok {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	// Mark complaint as resolved
	complaint.Resolved = true
	database.complaints[complaintID] = complaint

	// Return success message
	fmt.Fprint(w, "Complaint resolved successfully")
}

func handleResolveComplaint(w http.ResponseWriter, r *http.Request) {
	// Extract complaint ID from request
	complaintID := r.URL.Query().Get("complaint_id")

	// Retrieve complaint from database
	complaint, ok := database.complaints[complaintID]
	if !ok {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	// Mark complaint as resolved
	complaint.Resolved = true
	database.complaints[complaintID] = complaint

	// Return success message
	fmt.Fprint(w, "Complaint resolved successfully")
}

func handleResolveComplaint(w http.ResponseWriter, r *http.Request) {
	// Extract complaint ID from request
	complaintID := r.URL.Query().Get("complaint_id")

	// Retrieve complaint from database
	complaint, ok := database.complaints[complaintID]
	if !ok {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	// Mark complaint as resolved
	complaint.Resolved = true
	database.complaints[complaintID] = complaint

	// Return success message
	fmt.Fprint(w, "Complaint resolved successfully")
}

func handleResolveComplaint(w http.ResponseWriter, r *http.Request) {
	// Extract complaint ID from request
	complaintID := r.URL.Query().Get("complaint_id")

	// Retrieve complaint from database
	complaint, ok := database.complaints[complaintID]
	if !ok {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	// Mark complaint as resolved
	complaint.Resolved = true
	database.complaints[complaintID] = complaint

	// Return success message
	fmt.Fprint(w, "Complaint resolved successfully")
}

func generateID() string {
	// Generate unique ID using a UUID library or a similar method
	// For simplicity, we'll just return a placeholder string here
	return "123456"
}

func generateSecretCode() string {
	// Generate unique secret code using a UUID library or a similar method
	// For simplicity, we'll just return a placeholder string here
	return "abcdef"
}
