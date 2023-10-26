-- Create the Course table
CREATE TABLE course (
    id VARCHAR(10) PRIMARY KEY NOT NULL,
    lecturer_id VARCHAR(38) NOT NULL,
    name VARCHAR(255) NOT NULL,
    credit INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(38) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(38) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(38),
    FOREIGN KEY (lecturer_id) REFERENCES lecturer(id)
);
