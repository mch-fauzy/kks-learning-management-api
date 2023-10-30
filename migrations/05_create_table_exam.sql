-- Create the Exam table
CREATE TABLE exam (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    student_id VARCHAR(38) NOT NULL,
    course_id VARCHAR(38) NOT NULL,
    exam_type VARCHAR(50) NOT NULL,
    exam_date TIMESTAMP NOT NULL,
    score INT NOT NULL,
    letter_grade VARCHAR(2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(38) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(38) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(38),
    FOREIGN KEY (course_id) REFERENCES course(id),
    FOREIGN KEY (student_id) REFERENCES student(id)
);