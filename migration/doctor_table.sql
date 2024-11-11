-- Active: 1731001070739@@127.0.0.1@1433@o_clinic_DB


CREATE TABLE  doctor(
    doctor_id INT IDENTITY(1,1) PRIMARY KEY,    
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    specialization VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(25),
    hourly_rate DECIMAL(10,2),
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE()   
) 


CREATE TABLE working_days(
    doctor_id INT,
    day_of_week VARCHAR(255),
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE(),
    PRIMARY KEY (doctor_id, day_of_week),
    FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id) ON DELETE CASCADE
)

ALTER TABLE working_days ADD CONSTRAINT chk_day_of_week CHECK (day_of_week IN ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'))


