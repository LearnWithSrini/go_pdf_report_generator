package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LearnWithSrini/go_pdf_report_generator/models"
	"github.com/LearnWithSrini/go_pdf_report_generator/services"
)

// This demo shows PDF generation without requiring the backend to be running
func main() {
	fmt.Println("🎨 PDF Generation Demo")
	fmt.Println("======================")
	fmt.Println()

	// Create sample students
	students := []models.Student{
		{
			ID:          1,
			Name:        "John Doe",
			Email:       "john.doe@school.com",
			Phone:       "+1-234-567-8901",
			DOB:         "2005-03-15",
			FatherName:  "Robert Doe",
			FatherPhone: "+1-234-567-8900",
			MotherName:  "Mary Doe",
			MotherPhone: "+1-234-567-8902",
			Address:     "123 Main Street, Springfield, IL 62701",
			ClassName:   "Grade 10",
			SectionName: "A",
			Roll:        15,
			Status:      "active",
			CreatedAt:   time.Now().AddDate(0, -6, 0),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Jane Smith",
			Email:       "jane.smith@school.com",
			Phone:       "+1-234-567-8911",
			DOB:         "2006-07-22",
			FatherName:  "Michael Smith",
			FatherPhone: "+1-234-567-8910",
			MotherName:  "Sarah Smith",
			MotherPhone: "+1-234-567-8912",
			Address:     "456 Oak Avenue, Springfield, IL 62702",
			ClassName:   "Grade 9",
			SectionName: "B",
			Roll:        23,
			Status:      "active",
			CreatedAt:   time.Now().AddDate(0, -8, 0),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "Alex Johnson",
			Email:       "alex.johnson@school.com",
			Phone:       "+1-234-567-8921",
			DOB:         "2004-11-30",
			FatherName:  "David Johnson",
			FatherPhone: "+1-234-567-8920",
			MotherName:  "Lisa Johnson",
			MotherPhone: "+1-234-567-8922",
			Address:     "789 Pine Road, Springfield, IL 62703",
			ClassName:   "Grade 11",
			SectionName: "A",
			Roll:        8,
			Status:      "active",
			CreatedAt:   time.Now().AddDate(0, -10, 0),
			UpdatedAt:   time.Now(),
		},
	}

	// Create PDF generator
	pdfGen := services.NewPDFGenerator()

	// Generate PDFs for each student
	for _, student := range students {
		fmt.Printf("📄 Generating PDF for: %s (ID: %d)\n", student.Name, student.ID)

		pdfBytes, err := pdfGen.GenerateStudentReport(&student)
		if err != nil {
			log.Printf("❌ Error generating PDF for %s: %v", student.Name, err)
			continue
		}

		// Save PDF to file
		filename := fmt.Sprintf("demo_student_%d_report.pdf", student.ID)
		if err := savePDF(filename, pdfBytes); err != nil {
			log.Printf("❌ Error saving PDF: %v", err)
			continue
		}

		fmt.Printf("   ✅ Saved as: %s (%d bytes)\n", filename, len(pdfBytes))
		fmt.Println()
	}

	fmt.Println("🎉 Demo completed!")
	fmt.Println()
	fmt.Println("Generated PDFs:")
	for i := 1; i <= len(students); i++ {
		fmt.Printf("  - demo_student_%d_report.pdf\n", i)
	}
}

func savePDF(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
