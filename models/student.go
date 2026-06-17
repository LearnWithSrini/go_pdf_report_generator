package models

import "time"

// Student represents the student data structure from the backend API
type Student struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DOB         string    `json:"dob"`
	FatherName  string    `json:"father_name"`
	FatherPhone string    `json:"father_phone"`
	MotherName  string    `json:"mother_name"`
	MotherPhone string    `json:"mother_phone"`
	Address     string    `json:"address"`
	ClassName   string    `json:"class_name"`
	SectionName string    `json:"section_name"`
	Roll        int       `json:"roll"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// APIResponse represents the response structure from the backend
type APIResponse struct {
	Success bool    `json:"success"`
	Data    Student `json:"data"`
	Message string  `json:"message"`
}
