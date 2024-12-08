-- Active: 1733665822227@@db9994.public.databaseasp.net@1433@db9994


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

INSERT INTO doctor (first_name, last_name, specialization, email, phone, hourly_rate)
VALUES ('John', 'Doe', 'Cardiology', 'john.doe@example.com', '123-456-7890', 150.00);

INSERT INTO doctor (first_name, last_name, specialization, email, phone, hourly_rate)
VALUES ('Jane', 'Smith', 'Neurology', 'jane.smith@example.com', '234-567-8901', 200.00);

INSERT INTO doctor (first_name, last_name, specialization, email, phone, hourly_rate)
VALUES ('Emily', 'Johnson', 'Pediatrics', 'emily.johnson@example.com', '345-678-9012', 180.00);


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
    where d.doctor_id = 3
    GROUP BY d.doctor_id , d.first_name, d.last_name , d.specialization , d.email , d.phone , d.hourly_rate, d.created_at;
