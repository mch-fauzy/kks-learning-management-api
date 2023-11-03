-- Create the Student table
CREATE TABLE student (
    id VARCHAR(38) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    origin VARCHAR(100) NOT NULL,
    enrollment_date TIMESTAMP NOT NULL,
    gpa DECIMAL(3, 2),
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(38) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(38) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(38)
);