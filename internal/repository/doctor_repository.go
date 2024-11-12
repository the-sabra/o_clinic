package repository

import (
	"context"
	"database/sql"
	"o_clinic/internal/db"
	"o_clinic/internal/db/queries"
	"o_clinic/internal/domain"
	"strings"
)

// DoctorRepository represents the repository for the doctor.
type DoctorRepository struct {
	*db.Storage
}

// NewDoctorRepository creates a new doctor repository.
func NewDoctorRepository(storage *db.Storage) *DoctorRepository {
	return &DoctorRepository{Storage: storage}
}

// Create creates a new doctor.
func (r *DoctorRepository) Create(ctx context.Context, doctor *domain.NewDoctor) (int64, error) {
	var id int64
	err := r.DB.QueryRowContext(ctx, queries.CreateDoctor,
		doctor.FirstName,
		doctor.LastName,
		doctor.Specialization,
		doctor.Email,
		doctor.Phone,
		doctor.HourlyRate,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	for _, day := range doctor.WorkingDays {
		err := r.DB.QueryRowContext(ctx, queries.CreateDoctorWorkingDays, id, day).Err()
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

// GetByID retrieves a doctor by its ID.
func (r *DoctorRepository) GetByID(ctx context.Context, id int64) (*domain.Doctor, error) {
	doctor := &domain.Doctor{}
	err := r.DB.QueryRowContext(ctx, queries.GetDoctorByID, id).Scan(
		&doctor.ID,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.Specialization,
		&doctor.Email,
		&doctor.Phone,
		&doctor.HourlyRate,
		&doctor.WorkingDays,
		&doctor.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return doctor, nil
}

// GetPaginated retrieves a list of doctors with pagination.
func (r *DoctorRepository) GetPaginated(ctx context.Context, limit, offset int) ([]*domain.Doctor, error) {

	rows, err := r.DB.QueryContext(ctx, queries.ListDoctors, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	doctors := make([]*domain.Doctor, 0)
	for rows.Next() {
		doctor := &domain.Doctor{}
		var workingDays string
		err := rows.Scan(
			&doctor.ID,
			&doctor.FirstName,
			&doctor.LastName,
			&doctor.Specialization,
			&doctor.Email,
			&doctor.Phone,
			&doctor.HourlyRate,
			&doctor.CreatedAt,
			&workingDays,
		)
		if err != nil {
			return nil, err
		}
		doctor.WorkingDays = strings.Split(workingDays, ",")
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

// Update updates a doctor.
func (r *DoctorRepository) Update(ctx context.Context, doctor *domain.Doctor) error {
	_, err := r.DB.ExecContext(ctx, queries.UpdateDoctor,
		doctor.FirstName,
		doctor.LastName,
		doctor.Specialization,
		doctor.Email,
		doctor.Phone,
		doctor.HourlyRate,
		doctor.ID,
	)

	// delete all working days
	_, err = r.DB.ExecContext(ctx, queries.DeleteDoctorWorkingDays, doctor.ID)

	// insert new working days
	for _, day := range doctor.WorkingDays {
		_, err := r.DB.ExecContext(ctx, queries.CreateDoctorWorkingDays, doctor.ID, day)
		if err != nil {
			return err
		}
	}
	return err
}

// Delete deletes a doctor by its ID.
func (r *DoctorRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx, queries.DeleteDoctor, id)
	return err
}
