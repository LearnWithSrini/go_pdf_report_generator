package services

import (
	"bytes"
	"fmt"
	"time"

	"github.com/LearnWithSrini/go_pdf_report_generator/models"
	"github.com/jung-kurt/gofpdf"
)

type PDFGenerator struct{}

func NewPDFGenerator() *PDFGenerator {
	return &PDFGenerator{}
}

// GenerateStudentReport creates a PDF report for a student
func (g *PDFGenerator) GenerateStudentReport(student *models.Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set up colors
	headerColor := gofpdf.RGBType{R: 41, G: 128, B: 185}  // Blue
	sectionColor := gofpdf.RGBType{R: 52, G: 152, B: 219} // Light blue

	// Header
	pdf.SetFillColor(headerColor.R, headerColor.G, headerColor.B)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(0, 15, "Student Report", "", 1, "C", true, 0, "")
	pdf.Ln(5)

	// Reset text color
	pdf.SetTextColor(0, 0, 0)

	// Generated date
	pdf.SetFont("Arial", "I", 10)
	pdf.Cell(0, 6, "Generated on: "+time.Now().Format("January 2, 2006 at 3:04 PM"))
	pdf.Ln(10)

	// Personal Information Section
	g.addSectionHeader(pdf, "Personal Information", sectionColor)
	g.addField(pdf, "Student ID", fmt.Sprintf("%d", student.ID))
	g.addField(pdf, "Name", student.Name)
	g.addField(pdf, "Email", student.Email)
	g.addField(pdf, "Phone", student.Phone)
	g.addField(pdf, "Date of Birth", student.DOB)
	g.addField(pdf, "Status", student.Status)
	pdf.Ln(5)

	// Academic Information Section
	g.addSectionHeader(pdf, "Academic Information", sectionColor)
	g.addField(pdf, "Class", student.ClassName)
	g.addField(pdf, "Section", student.SectionName)
	g.addField(pdf, "Roll Number", fmt.Sprintf("%d", student.Roll))
	pdf.Ln(5)

	// Parent Information Section
	g.addSectionHeader(pdf, "Parent/Guardian Information", sectionColor)
	if student.FatherName != "" {
		g.addField(pdf, "Father's Name", student.FatherName)
		g.addField(pdf, "Father's Phone", student.FatherPhone)
	}
	if student.MotherName != "" {
		g.addField(pdf, "Mother's Name", student.MotherName)
		g.addField(pdf, "Mother's Phone", student.MotherPhone)
	}
	pdf.Ln(5)

	// Address Section
	if student.Address != "" {
		g.addSectionHeader(pdf, "Address", sectionColor)
		pdf.SetFont("Arial", "", 11)
		pdf.MultiCell(0, 6, student.Address, "", "L", false)
		pdf.Ln(5)
	}

	// Footer
	pdf.SetY(-20)
	pdf.SetFont("Arial", "I", 8)
	pdf.SetTextColor(128, 128, 128)
	pdf.CellFormat(0, 10, "This is a computer-generated report. No signature required.", "", 0, "C", false, 0, "")

	// Get PDF bytes
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %v", err)
	}

	return buf.Bytes(), nil
}

func (g *PDFGenerator) addSectionHeader(pdf *gofpdf.Fpdf, title string, color gofpdf.RGBType) {
	pdf.SetFillColor(color.R, color.G, color.B)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, title, "", 1, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(2)
}

func (g *PDFGenerator) addField(pdf *gofpdf.Fpdf, label, value string) {
	if value == "" || value == "0" {
		return
	}

	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(50, 8, label+":")
	pdf.SetFont("Arial", "", 11)
	pdf.Cell(0, 8, value)
	pdf.Ln(8)
}
