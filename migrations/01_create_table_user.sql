-- Create the User table
CREATE TABLE user (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    role ENUM('admin', 'student', 'lecturer') NOT NULL,
    last_login TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(36) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(36) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(36),
    CONSTRAINT valid_email CHECK (email REGEXP '^[A-Za-z0-9._]+@[A-Za-z.]+\\.[A-Za-z]{2,4}$')
);