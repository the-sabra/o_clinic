package queries


const (
	// CreatePatient is the SQL query to insert a new patient and return the patient_id.
	CreatePatient = `
		INSERT INTO patient (first_name, last_name, dob, gender, email, phone)
		OUTPUT INSERTED.patient_id
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6)`

	// GetPatientByID is the SQL query to select a patient by patient_id.
	GetPatientByID = `
		SELECT patient_id, first_name, last_name, dob, gender, email, phone
		FROM patient
		WHERE patient_id = @p1`

	// GetPatientByEmail is the SQL query to select a patient by email.
	GetPatientByEmail = `
		SELECT patient_id, first_name, last_name, dob , gender, email, phone
		FROM patient
		WHERE email = @p1`
	
	// UpdatePatient is the SQL query to update a patient by patient_id.
	UpdatePatient = `
		UPDATE patient
		SET first_name = @p1, 
		last_name = @p2, 
		dob = @p3 ,
		email = @p4, 
		phone = @p5
		WHERE patient_id = @p6`

	// ListPatients is the SQL query to select all patients
	ListPatients = `
  		SELECT
			p.patient_id,
			p.first_name,
			p.last_name,
			p.dob,
			p.gender,
			p.email,
			p.phone
		FROM patient p
		ORDER BY p.created_at
		OFFSET 0 ROWS FETCH NEXT 5 ROWS ONLY;`
	
	// DeletePatient is the SQL query to delete a patient by patient_id.
	DeletePatient = `
		DELETE FROM patient
		WHERE patient_id = @p1`
) 