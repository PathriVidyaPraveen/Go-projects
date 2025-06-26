package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Submission struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

var (
	submissions []Submission
	mu          sync.Mutex
	submitChan  = make(chan Submission)
)

func main() {
	fmt.Println("GoContact API server starting...")
	http.HandleFunc("/submit", handleSubmit)
	http.HandleFunc("/submissions", handleGetSubmissions)

	go processSubmissions()

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var sub Submission
	err := json.NewDecoder(r.Body).Decode(&sub)
	if err != nil || sub.Name == "" || sub.Email == "" || sub.Message == "" || !strings.Contains(sub.Email, "@") {
		http.Error(w, "Invalid submission!", http.StatusBadRequest)
		return
	}

	submitChan <- sub
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Submission received")
}

func handleGetSubmissions(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submissions)
}

func processSubmissions() {
	for sub := range submitChan {
		mu.Lock()
		submissions = append(submissions, sub)
		mu.Unlock()
		fmt.Printf("Processed: %s <%s>\n", sub.Name, sub.Email)
	}
}
