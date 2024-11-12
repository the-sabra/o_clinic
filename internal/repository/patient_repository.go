package repository

import (
	"context"
	"fmt"
	"o_clinic/internal/db"
	"o_clinic/internal/db/queries"
	"o_clinic/internal/domain"
)

// PatientRepository represents the repository for the patient.
type PatientRepository struct {
	*db.Storage
}

// NewPatientRepository creates a new patient repository.
func NewPatientRepository(storage *db.Storage) *PatientRepository {
	return &PatientRepository{Storage: storage}
}

// Create creates a new patient.
func (r *PatientRepository) Create(ctx context.Context, patient *domain.NewPatient) (int64, error) {
	var id int64
	err := r.DB.QueryRowContext(ctx, queries.CreatePatient,
		patient.FirstName,
		patient.LastName,
		patient.DateOfBirth,
		patient.Gender,
		patient.Email,
		patient.Phone,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetByID retrieves a patient by its ID.
func(r *PatientRepository) GetByID(ctx context.Context , id int64) (*domain.Patient, error) {
	patient := &domain.Patient{}
	err := r.DB.QueryRowContext(ctx, queries.GetPatientByID, id).Scan(
		&patient.ID,
		&patient.FirstName,
		&patient.LastName,
		&patient.DateOfBirth,
		&patient.Gender,
		&patient.Email,
		&patient.Phone,
	)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

// GetPaginated retrieves a list of patients with pagination.
func(r *PatientRepository) GetPaginated(ctx context.Context , limit, offset int) ([]*domain.Patient, error) {
	rows, err := r.DB.QueryContext(ctx, queries.ListPatients, offset, limit)
	if err != nil {
		return nil, err
	}
 
	defer rows.Close()

	patients := make([]*domain.Patient, 0)
	for rows.Next() {
		patient := &domain.Patient{}
		err := rows.Scan(
			&patient.ID,
			&patient.FirstName,
			&patient.LastName,
			&patient.DateOfBirth,
			&patient.Gender,
			&patient.Email,
			&patient.Phone,
		)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

// Update updates a patient.
func(r *PatientRepository) Update(ctx context.Context , patient *domain.Patient) error {
	rows, err := r.DB.ExecContext(ctx, queries.UpdatePatient,
		patient.FirstName,
		patient.LastName,
		patient.DateOfBirth,
		patient.Email,
		patient.Phone,
		patient.ID,
	)

	// check if the patient exists
	if rows, _ := rows.RowsAffected(); rows == 0 {
		return fmt.Errorf("patient with ID %d does not exist", patient.ID)
	}

    return err
} 

// Delete deletes a patient.
func(r *PatientRepository) Delete(ctx context.Context , id int64) error {
	_, err := r.DB.ExecContext(ctx, queries.DeletePatient, id)
	return err
}