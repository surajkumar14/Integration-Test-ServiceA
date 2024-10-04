-- Create the 'dummy_users' table
CREATE TABLE IF NOT EXISTS dummy_users (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    age INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data into the 'dummy_users' table
INSERT INTO dummy_users (id, name, email, age) VALUES (1, 'John Doe', 'john.doe@example.com', 30);
INSERT INTO dummy_users (id, name, email, age) VALUES (2, 'Jane Doe', 'jane.doe@example.com', 28);
INSERT INTO dummy_users (id, name, email, age) VALUES (3, 'Alice Smith', 'alice.smith@example.com', 35);
INSERT INTO dummy_users (id, name, email, age) VALUES (4, 'Bob Johnson', 'bob.johnson@example.com', 40);
INSERT INTO dummy_users (id, name, email, age) VALUES (5, 'Charlie Brown', 'charlie.brown@example.com', 25);
INSERT INTO dummy_users (id, name, email, age) VALUES (6, 'David Wilson', 'david.wilson@example.com', 45);
INSERT INTO dummy_users (id, name, email, age) VALUES (7, 'Eve Davis', 'eve.davis@example.com', 32);
INSERT INTO dummy_users (id, name, email, age) VALUES (8, 'Frank Miller', 'frank.miller@example.com', 29);
INSERT INTO dummy_users (id, name, email, age) VALUES (9, 'Grace Lee', 'grace.lee@example.com', 27);
INSERT INTO dummy_users (id, name, email, age) VALUES (10, 'Hank Green', 'hank.green@example.com', 33);
INSERT INTO dummy_users (id, name, email, age) VALUES (11, 'Ivy White', 'ivy.white@example.com', 31);
INSERT INTO dummy_users (id, name, email, age) VALUES (12, 'Jack Black', 'jack.black@example.com', 38);
INSERT INTO dummy_users (id, name, email, age) VALUES (13, 'Karen Scott', 'karen.scott@example.com', 26);
INSERT INTO dummy_users (id, name, email, age) VALUES (14, 'Leo King', 'leo.king@example.com', 34);
INSERT INTO dummy_users (id, name, email, age) VALUES (15, 'Mona Taylor', 'mona.taylor@example.com', 37);
INSERT INTO dummy_users (id, name, email, age) VALUES (16, 'Nina Brown', 'nina.brown@example.com', 24);
INSERT INTO dummy_users (id, name, email, age) VALUES (17, 'Oscar Green', 'oscar.green@example.com', 39);
INSERT INTO dummy_users (id, name, email, age) VALUES (18, 'Paul Adams', 'paul.adams@example.com', 36);
INSERT INTO dummy_users (id, name, email, age) VALUES (19, 'Quincy Hall', 'quincy.hall@example.com', 41);
INSERT INTO dummy_users (id, name, email, age) VALUES (20, 'Rachel Young', 'rachel.young@example.com', 22);