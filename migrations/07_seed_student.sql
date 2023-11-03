-- Insert dummy data into the Student table
INSERT INTO student (id, name, origin, enrollment_date, gpa, created_at, created_by, updated_at, updated_by)
VALUES
    ('S-MOCK1', 'Aditya Pamungkas', 'Magelang', '2023-01-15 10:00:00', 3.75, CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('S-MOCK2', 'Adi Mahendra', 'Surabaya', '2023-01-20 11:30:00', 3.90, CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('S-MOCK3', 'Hendra Suryo', 'Bekasi', '2023-02-05 09:15:00', 3.45, CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin');
