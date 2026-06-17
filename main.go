package main

import (
	"log"
	"net/http"
	"os"

	"github.com/LearnWithSrini/go_pdf_report_generator/config"
	"github.com/LearnWithSrini/go_pdf_report_generator/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize router
	router := mux.NewRouter()

	// Initialize handlers
	studentHandler := handlers.NewStudentHandler(cfg)

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/students/{id}/report", studentHandler.GenerateReport).Methods("GET")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}

	log.Printf("Go PDF Service starting on port %s", port)
	log.Printf("Backend API URL: %s", cfg.BackendAPIURL)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
