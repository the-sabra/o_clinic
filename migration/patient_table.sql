

CREATE TABLE patient (
    patient_id INT IDENTITY(1,1) PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    dob DATE,
    gender VARCHAR(10),
    email VARCHAR(255),
    phone VARCHAR(25),
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE()
);  

DROP TABLE patient;

--make bulk insertion of patients
INSERT INTO patient (first_name, last_name, dob, gender, email, phone)
VALUES
('John', 'Doe', '1999-01-01', 'MALE', 'test@mail.com','1234567890'),
('Jane', 'Smith', '1985-05-15', 'FEMALE', 'jane.smith@mail.com', '0987654321'),
('Alice', 'Johnson', '1978-09-23', 'FEMALE', 'alice.johnson@mail.com', '1122334455'),
('Bob', 'Brown', '1990-12-11', 'MALE', 'bob.brown@mail.com', '2233445566'),
('Charlie', 'Davis', '2000-07-07', 'MALE', 'charlie.davis@mail.com', '3344556677'),
('Eve', 'Miller', '1995-03-30', 'FEMALE', 'eve.miller@mail.com', '4455667788'),
('Frank', 'Wilson', '1988-11-22', 'MALE', 'frank.wilson@mail.com', '5566778899'),
('Grace', 'Lee', '1992-04-18', 'FEMALE', 'grace.lee@mail.com', '6677889900'),
('Hank', 'Taylor', '1983-08-30', 'MALE', 'hank.taylor@mail.com', '7788990011'),
('Ivy', 'Anderson', '1975-02-14', 'FEMALE', 'ivy.anderson@mail.com', '8899001122'),
('Jack', 'Thomas', '1997-06-25', 'MALE', 'jack.thomas@mail.com', '9900112233'),
('Karen', 'White', '1982-10-10', 'FEMALE', 'karen.white@mail.com', '1011121314'),
('Leo', 'Harris', '1991-01-20', 'MALE', 'leo.harris@mail.com', '1213141516'),
('Mona', 'Clark', '1987-05-05', 'FEMALE', 'mona.clark@mail.com', '1314151617'),
('Nate', 'Lewis', '1993-03-03', 'MALE', 'nate.lewis@mail.com', '1415161718'),
('Olivia', 'Walker', '1980-12-12', 'FEMALE', 'olivia.walker@mail.com', '1516171819'),
('Paul', 'Hall', '1996-07-07', 'MALE', 'paul.hall@mail.com', '1617181920'),
('Quinn', 'Allen', '1984-09-09', 'FEMALE', 'quinn.allen@mail.com', '1718192021'),
('Rita', 'Young', '1994-11-11', 'FEMALE', 'rita.young@mail.com', '1819202122'),
('Sam', 'King', '1981-02-02', 'MALE', 'sam.king@mail.com', '1920212223'),
('Tina', 'Scott', '1989-04-04', 'FEMALE', 'tina.scott@mail.com', '2021222324'),
('Uma', 'Nelson', '1990-05-05', 'FEMALE', 'uma.nelson@mail.com', '2122232425'),
('Victor', 'Perez', '1986-06-06', 'MALE', 'victor.perez@mail.com', '2223242526'),
('Wendy', 'Carter', '1992-07-07', 'FEMALE', 'wendy.carter@mail.com', '2324252627'),
('Xander', 'Mitchell', '1984-08-08', 'MALE', 'xander.mitchell@mail.com', '2425262728'),
('Yara', 'Roberts', '1991-09-09', 'FEMALE', 'yara.roberts@mail.com', '2526272829'),
('Zane', 'Evans', '1983-10-10', 'MALE', 'zane.evans@mail.com', '2627282930'),
('Amy', 'Green', '1995-11-11', 'FEMALE', 'amy.green@mail.com', '2728293031'),
('Brian', 'Baker', '1987-12-12', 'MALE', 'brian.baker@mail.com', '2829303132'),
('Cathy', 'Adams', '1990-01-01', 'FEMALE', 'cathy.adams@mail.com', '2930313233'),
('David', 'Clark', '1985-02-02', 'MALE', 'david.clark@mail.com', '3031323334'),
('Ella', 'Lewis', '1993-03-03', 'FEMALE', 'ella.lewis@mail.com', '3132333435'),
('Fred', 'Walker', '1988-04-04', 'MALE', 'fred.walker@mail.com', '3233343536'),
('Gina', 'Hall', '1992-05-05', 'FEMALE', 'gina.hall@mail.com', '3334353637'),
('Harry', 'Allen', '1984-06-06', 'MALE', 'harry.allen@mail.com', '3435363738'),
('Iris', 'Young', '1991-07-07', 'FEMALE', 'iris.young@mail.com', '3536373839'),
('Jake', 'King', '1983-08-08', 'MALE', 'jake.king@mail.com', '3637383940'),
('Kara', 'Scott', '1989-09-09', 'FEMALE', 'kara.scott@mail.com', '3738394041'),
('Liam', 'Nelson', '1990-10-10', 'MALE', 'liam.nelson@mail.com', '3839404142'),
('Mia', 'Perez', '1986-11-11', 'FEMALE', 'mia.perez@mail.com', '3940414243'),
('Noah', 'Carter', '1992-12-12', 'MALE', 'noah.carter@mail.com', '4041424344'),
('Olga', 'Mitchell', '1984-01-01', 'FEMALE', 'olga.mitchell@mail.com', '4142434445'),
('Pete', 'Roberts', '1991-02-02', 'MALE', 'pete.roberts@mail.com', '4243444546'),
('Quincy', 'Evans', '1983-03-03', 'MALE', 'quincy.evans@mail.com', '4344454647'),
('Rachel', 'Green', '1995-04-04', 'FEMALE', 'rachel.green@mail.com', '4445464748'),
('Steve', 'Baker', '1987-05-05', 'MALE', 'steve.baker@mail.com', '4546474849'),
('Tina', 'Adams', '1990-06-06', 'FEMALE', 'tina.adams@mail.com', '4647484950');


