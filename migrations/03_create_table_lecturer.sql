-- Create the Lecturer table
CREATE TABLE lecturer (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    user_id VARCHAR(36) NOT NULL UNIQUE,
    registration_number BIGINT(15) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    origin VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(36) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(36) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(36),
    FOREIGN KEY (user_id) REFERENCES user(id)
);