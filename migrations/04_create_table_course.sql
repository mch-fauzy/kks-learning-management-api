-- Create the Course table
CREATE TABLE course (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    lecturer_id VARCHAR(36),
    name VARCHAR(255) NOT NULL,
    credit INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(36) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(36) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(36),
    FOREIGN KEY (lecturer_id) REFERENCES lecturer(id)
);
