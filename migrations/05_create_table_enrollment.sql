-- Create the Enrollment table
CREATE TABLE enrollment (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    student_id VARCHAR(36) NOT NULL,
    course_id VARCHAR(36) NOT NULL,
    course_enrollment_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(36) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(36) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(36),
    FOREIGN KEY (student_id) REFERENCES student(id),
    FOREIGN KEY (course_id) REFERENCES course(id)
);