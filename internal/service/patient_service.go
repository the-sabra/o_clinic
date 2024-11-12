package service

import (
	"context"
	"o_clinic/internal/domain"
)

// PatientRepository represents the repository for the patient.
type PatientRepository interface {
	Create(ctx context.Context, patient *domain.NewPatient) (int64, error)
	GetByID(ctx context.Context , id int64) (*domain.Patient, error)
	GetPaginated(ctx context.Context , limit, offset int) ([]*domain.Patient, error)
	Update(ctx context.Context , patient *domain.Patient) error
	Delete(ctx context.Context , id int64) error 
}

// PatientService represents the service for the patient.
type PatientService struct {
	repo PatientRepository
}


// NewPatientService creates a new patient service.
func NewPatientService(repo PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}


// CreatePatient creates a new patient.
func (s *PatientService) CreatePatient(ctx context.Context, patient *domain.NewPatient) (int64, error) {
	// Add any business logic here
	return s.repo.Create(ctx, patient)
}

// GetPatient retrieves a patient by its ID.
func (s *PatientService) GetPatient(ctx context.Context, id int64) (*domain.Patient, error) {
	return s.repo.GetByID(ctx, id)
}

// GetAllPagination retrieves all patients with pagination.
func (s *PatientService) GetAllPagination(ctx context.Context, limit, offset int) ([]*domain.Patient, error) {
	return s.repo.GetPaginated(ctx, limit, offset)
}

// UpdatePatient updates a patient.
func (s *PatientService) UpdatePatient(ctx context.Context, patient *domain.Patient) error {
	return s.repo.Update(ctx, patient)
}

// DeletePatient deletes a patient by its ID.
func (s *PatientService) DeletePatient(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx,id)
} 


