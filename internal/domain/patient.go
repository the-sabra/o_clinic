package domain

// Patient represents a patient with various attributes.
type Patient struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CreatedAt   string `json:"created_at,omitempty"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"dob"`
}

// NewPatient represents a new patient to be created.
type NewPatient struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}
