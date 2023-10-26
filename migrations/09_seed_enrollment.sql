-- Insert dummy data into the Enrollment table
INSERT INTO enrollment (student_id, course_id, course_enrollment_date, created_at, created_by, updated_at, updated_by)
VALUES
    ('S-MOCK1', 'C-MOCK1', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('S-MOCK1', 'C-MOCK3', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('S-MOCK2', 'C-MOCK3', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('S-MOCK3', 'C-MOCK2', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('S-MOCK3', 'C-MOCK3', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin');