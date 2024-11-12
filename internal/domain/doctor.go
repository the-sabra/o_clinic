package domain

// Specialization represents the specialization of a doctor.

// Doctor represents a medical doctor with various attributes.
import (
	"errors"
	"time"
)

// Specialization represents the specialization of a doctor.
type Specialization string

// Doctor represents a medical doctor with various attributes.
type Doctor struct {
	ID             int64          `json:"id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Specialization Specialization `json:"specialization"`
	Email          string         `json:"email"`
	Phone          *string        `json:"phone,omitempty"` //
	HourlyRate     float64        `json:"hourly_rate"`
	WorkingDays    []string       `json:"working_days"`
	CreatedAt      time.Time      `json:"created_at,omitempty"`
}

// NewDoctor represents a new doctor to be created.
type NewDoctor struct {
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Specialization Specialization `json:"specialization"`
	Email          string         `json:"email"`
	Phone          *string        `json:"phone,omitempty"`
	HourlyRate     float64        `json:"hourly_rate"`
	WorkingDays    []string       `json:"working_days"`
}

// WorkingDays represents the working days of a doctor.
type WorkingDays struct {
	doctorID int64
	day      string
}

// validate checks if the doctor struct is valid.
func (d *Doctor) validate() error {
	if d.FirstName == "" {
		return errors.New("first name is required")
	}
	if d.LastName == "" {
		return errors.New("last name is required")
	}
	if d.Specialization == "" {
		return errors.New("specialization is required")
	}
	if d.Email == "" {
		return errors.New("email is required")
	}
	if d.HourlyRate == 0 {
		return errors.New("hourly rate is required")
	}
	if len(d.WorkingDays) == 0 {
		return errors.New("working days is required")
	}
	return nil
}
