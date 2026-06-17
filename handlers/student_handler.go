package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LearnWithSrini/go_pdf_report_generator/config"
	"github.com/LearnWithSrini/go_pdf_report_generator/services"
	"github.com/gorilla/mux"
)

type StudentHandler struct {
	backendClient *services.BackendClient
	pdfGenerator  *services.PDFGenerator
}

func NewStudentHandler(cfg *config.Config) *StudentHandler {
	return &StudentHandler{
		backendClient: services.NewBackendClient(cfg.BackendAPIURL, cfg.BackendTimeout),
		pdfGenerator:  services.NewPDFGenerator(),
	}
}

// GenerateReport handles GET /api/v1/students/:id/report
func (h *StudentHandler) GenerateReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	if studentID == "" {
		h.sendError(w, "Student ID is required", http.StatusBadRequest)
		return
	}

	log.Printf("Generating report for student ID: %s", studentID)

	// Fetch student data from backend API
	student, err := h.backendClient.GetStudent(studentID)
	if err != nil {
		log.Printf("Error fetching student data: %v", err)
		h.sendError(w, fmt.Sprintf("Failed to fetch student data: %v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully fetched student data: %s", student.Name)

	// Generate PDF
	pdfBytes, err := h.pdfGenerator.GenerateStudentReport(student)
	if err != nil {
		log.Printf("Error generating PDF: %v", err)
		h.sendError(w, fmt.Sprintf("Failed to generate PDF: %v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully generated PDF report for student: %s", student.Name)

	// Set response headers
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=student_%s_report.pdf", studentID))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))

	// Write PDF to response
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}

func (h *StudentHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"success": false, "error": "%s"}`, message)))
}
