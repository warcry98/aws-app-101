CREATE TABLE users (
    user_id int PRIMARY KEY,
    username varchar(25) NOT NULL,
    fullname varchar(255),
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL
);

INSERT INTO users (
    user_id,
    username,
    email,
    password
) VALUES (
    1,
    'Tester-1',
    'Tester@gmail.com',
    '$2a$10$SncukdqYaeNqiTOs9a9pGeIjLKhZ8lMEkm7efOXPcnkbnsn7rSjrG'
);
