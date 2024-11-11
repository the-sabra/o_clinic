package queries

const (
    // CreateDoctor is the SQL query to insert a new doctor and return the doctor_id.
    CreateDoctor = `
        INSERT INTO doctor (first_name, last_name, specialization, email, phone, hourly_rate)
        OUTPUT INSERTED.doctor_id
        VALUES (@p1, @p2, @p3, @p4, @p5, @p6)`

    // CreateDoctorWorkingDays is the SQL query to insert the working days of a doctor.
    CreateDoctorWorkingDays = `
        INSERT INTO working_days (doctor_id, day_of_week)
        VALUES (@p1, @p2)`
    
    // GetDoctorByID is the SQL query to select a doctor by doctor_id.
    GetDoctorByID = `
        SELECT doctor_id, first_name, last_name, specialization, email, phone, hourly_rate, working_days, created_at
        FROM doctor
        WHERE doctor_id = @p1`

    // ListDoctors is the SQL query to select all doctors.
    ListDoctors = `
    SELECT 
        d.doctor_id, 
        d.first_name, 
        d.last_name, 
        d.specialization, 
        d.email, 
        d.phone, 
        d.hourly_rate,
        d.created_at,       
        STRING_AGG(w.day_of_week, ',') AS working_days
    FROM doctor d
    LEFT JOIN working_days w ON d.doctor_id = w.doctor_id
    GROUP BY d.doctor_id , d.first_name, d.last_name , d.specialization , d.email , d.phone , d.hourly_rate, d.created_at
    ORDER BY d.created_at
    OFFSET @p1 ROWS FETCH NEXT @p2 ROWS ONLY;`

    // UpdateDoctor is the SQL query to update a doctor by doctor_id.
    UpdateDoctor = `
    UPDATE doctor 
    SET first_name = @p1, last_name = @p2, specialization = @p3, email = @p4, phone = @p5, hourly_rate = @p6
    WHERE doctor_id = @p7`

    // DeleteDoctorWorkingDays is the SQL query to delete all working days of a doctor.
    DeleteDoctorWorkingDays = `
    DELETE FROM working_days
    WHERE doctor_id = @p1`

    // DeleteDoctor is the SQL query to delete a doctor by doctor_id.
    DeleteDoctor = `
    DELETE FROM doctor
    WHERE doctor_id = @p1`

)     