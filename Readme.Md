 CREATE TABLE attendance (
    attendance_id INT PRIMARY KEY AUTO_INCREMENT,
    date DATE NOT NULL,
    class_id INT NOT NULL,
    student_id INT NOT NULL,
    status VARCHAR(10) NOT NULL,
    subject VARCHAR(100),
    teacher_id INT,
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (class_id) REFERENCES classes(class_id),
    FOREIGN KEY (teacher_id) REFERENCES staff(staff_id)
);


 results
CREATE TABLE exam_results (
    exam_result_id INT PRIMARY KEY AUTO_INCREMENT,
    exam_id INT NOT NULL,
    student_id INT NOT NULL,
    subject VARCHAR(100) NOT NULL,
    marks_obtained DECIMAL(5,2) NOT NULL,
    total_marks DECIMAL(5,2) NOT NULL,
    grade VARCHAR(2),
    percentage DECIMAL(5,2),
    status VARCHAR(20),
    FOREIGN KEY (exam_id) REFERENCES exams(exam_id),
    FOREIGN KEY (student_id) REFERENCES students(student_id)
);


 homework
 CREATE TABLE homework (
    homework_id INT PRIMARY KEY AUTO_INCREMENT,
    teacher_id INT NOT NULL,
    class_id INT NOT NULL,
    subject VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    due_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL,
    student_id INT,
    submission_date DATE,
    grade VARCHAR(5),
    FOREIGN KEY (teacher_id) REFERENCES staff(staff_id),
    FOREIGN KEY (class_id) REFERENCES classes(class_id),
    FOREIGN KEY (student_id) REFERENCES students(student_id)
);


 notifications

CREATE TABLE notes_and_learning (
    note_id INT PRIMARY KEY AUTO_INCREMENT,
    teacher_id INT NOT NULL,
    class_id INT NOT NULL,
    subject VARCHAR(100) NOT NULL,
    note_type VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    file_path VARCHAR(255),
    upload_date DATE NOT NULL,
    last_updated DATE,
    status VARCHAR(20) NOT NULL,
    FOREIGN KEY (teacher_id) REFERENCES staff(staff_id),
    FOREIGN KEY (class_id) REFERENCES classes(class_id)
);

CREATE TABLE notifications (
    notification_id INT PRIMARY KEY AUTO_INCREMENT,
    sender_id INT NOT NULL,
    recipient_id INT NOT NULL,
    notification_type VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    read_at TIMESTAMP NULL,
    priority VARCHAR(10),
    event_date DATE,
    FOREIGN KEY (sender_id) REFERENCES staff(staff_id),
    FOREIGN KEY (recipient_id) REFERENCES students(student_id)
);



 daily time table
 CREATE TABLE daily_timetable (
    time_slot_id INT PRIMARY KEY AUTO_INCREMENT,
    day_of_week VARCHAR(10) NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    subject VARCHAR(100) NOT NULL,
    class_id INT NOT NULL,
    teacher_id INT,
    FOREIGN KEY (teacher_id) REFERENCES staff(staff_id),
    FOREIGN KEY (class_id) REFERENCES classes(class_id)
);
  
chat microservices
post microservices
websites microservices
attachment microservices

