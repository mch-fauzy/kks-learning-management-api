-- Insert dummy data into the Lecturer table
INSERT INTO lecturer (id, name, origin, created_at, created_by, updated_at, updated_by)
VALUES
    ('L-MOCK1', 'Diky Anggoro', 'Bandung', CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin'),
    ('L-MOCK2', 'Putri Ayu', 'Surabaya', CURRENT_TIMESTAMP(), 'Admin', CURRENT_TIMESTAMP(), 'Admin');