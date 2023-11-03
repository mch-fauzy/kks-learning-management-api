-- Create the Exam table
CREATE TABLE exam (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    student_id VARCHAR(36) NOT NULL,
    course_id VARCHAR(36) NOT NULL,
    exam_type VARCHAR(50),
    exam_date TIMESTAMP,
    score INT NOT NULL,
    letter_grade VARCHAR(2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(36) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(36) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(36),
    FOREIGN KEY (course_id) REFERENCES course(id),
    FOREIGN KEY (student_id) REFERENCES student(id)
);