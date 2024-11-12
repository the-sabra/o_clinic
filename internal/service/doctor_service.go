package service

import (
	"context"
	"o_clinic/internal/domain"
)


// DoctorRepository represents the repository for the doctor.
type DoctorRepository interface {
	Create(ctx context.Context, doctor *domain.NewDoctor) (int64, error)
	GetByID(ctx context.Context, id int64) (*domain.Doctor, error)
	GetPaginated(ctx context.Context, limit, offset int) ([]*domain.Doctor, error)
	Update(ctx context.Context, doctor *domain.Doctor) error
	Delete(ctx context.Context, id int64) error
}


// DoctorService represents the service for the doctor.
type DoctorService struct {
	repo DoctorRepository
}

// NewDoctorService creates a new doctor service.
func NewDoctorService(repo DoctorRepository) *DoctorService {
	return &DoctorService{repo: repo}
}

// CreateDoctor creates a new doctor.
func (s *DoctorService) CreateDoctor(ctx context.Context, doctor *domain.NewDoctor) (int64, error) {
	// Add any business logic here
	return s.repo.Create(ctx, doctor)
}

// GetDoctor retrieves a doctor by its ID.
func (s *DoctorService) GetDoctor(ctx context.Context, id int64) (*domain.Doctor, error) {
	return s.repo.GetByID(ctx, id)
}

// GetAllPagination retrieves all doctors with pagination.
func (s *DoctorService) GetAllPagination(ctx context.Context, limit, offset int) ([]*domain.Doctor, error) {
	return s.repo.GetPaginated(ctx, limit, offset)
}

// UpdateDoctor updates a doctor.
func (s *DoctorService) UpdateDoctor(ctx context.Context, doctor *domain.Doctor) error {
	return s.repo.Update(ctx, doctor)
}

// DeleteDoctor deletes a doctor by its ID.
func (s *DoctorService) DeleteDoctor(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
