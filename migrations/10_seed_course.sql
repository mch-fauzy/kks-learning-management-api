-- TODO
-- Insert dummy data into the Course table
INSERT INTO course (id, lecturer_id, name, credit, created_at, created_by, updated_at, updated_by)
VALUES
    ('C-MOCK1', 'L-MOCK1', 'Biologi A', 3, CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('C-MOCK2', 'L-MOCK1', 'Biologi B', 3, CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('C-MOCK3', 'L-MOCK2', 'Statistika A', 4, CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin');