-- Insert multiple sample users with UUIDs and hashed passwords using Bcrypt
INSERT INTO user (id, email, password, role, created_at, created_by, updated_at, updated_by)
VALUES
('a8297546-88d2-4e53-8cc1-9a9d5c961f4f', 'admin.example@admin.kks.ac.id', '$2a$10$GPyygUTIW5eBzCZ0nskHc.WhzrHjWOhtWFMxDK.wRYZMOuMUb0orO', 'admin', NOW(), 'a8297546-88d2-4e53-8cc1-9a9d5c961f4f', NOW(), 'a8297546-88d2-4e53-8cc1-9a9d5c961f4f'),
('b57a5c8b-c77e-4e6b-a9b2-25f42fc6a6f3', 'student.example@student.kks.ac.id', '$2a$10$nAuwk7haLLn9lpA.VX7ml.1g0VDs71SpM.a7KhL2rME.1IRFa1ou.', 'student', NOW(), 'a8297546-88d2-4e53-8cc1-9a9d5c961f4f', NOW(), 'a8297546-88d2-4e53-8cc1-9a9d5c961f4f'),
('c26a85de-e441-47a4-9ca3-0c5d3abf1e32', 'lecturer.example@lecturer.kks.ac.id', '$2a$10$8ysM2ZQ69Fg/P54WGt0vSOeGGEsrm0B41E99Pt3xdpLHIjqEZVLJS', 'lecturer', NOW(), 'a8297546-88d2-4e53-8cc1-9a9d5c961f4f', NOW(), 'a8297546-88d2-4e53-8cc1-9a9d5c961f4f');

-- Email | Password
-- admin.example@admin.kks.ac.id | admin_password
-- student.example@student.kks.ac.id | student_password
-- lecturer.example@lecturer.kks.ac.id | lecturer_password
