package service

import (
	"context"
	"o_clinic/internal/domain"
)


type DoctorRepository interface {
    Create(ctx context.Context, doctor *domain.NewDoctor) (int64, error)
    GetByID(ctx context.Context, id int64) (*domain.Doctor, error)
    GetPaginated(ctx context.Context, limit, offset int) ([]*domain.Doctor, error) 
    Update(ctx context.Context, doctor *domain.Doctor) error
    Delete(ctx context.Context, id int64) error
}

type DoctorService struct {
    repo DoctorRepository
}

func NewDoctorService(repo DoctorRepository) *DoctorService {
    return &DoctorService{repo: repo}
}

func (s *DoctorService) CreateDoctor(ctx context.Context, doctor *domain.NewDoctor) (int64, error) {
    // Add any business logic here
    return s.repo.Create(ctx, doctor)
}

func (s *DoctorService) GetDoctor(ctx context.Context, id int64) (*domain.Doctor, error) {
    return s.repo.GetByID(ctx, id)
} 

func (s *DoctorService) GetAllPagination(ctx context.Context, limit, offset int)([]*domain.Doctor, error) {
    return s.repo.GetPaginated(ctx, limit, offset) 
} 

func (s *DoctorService) UpdateDoctor(ctx context.Context, doctor *domain.Doctor) error {
    return s.repo.Update(ctx, doctor)
}

func (s *DoctorService) DeleteDoctor(ctx context.Context, id int64) error {
    return s.repo.Delete(ctx, id)
}
