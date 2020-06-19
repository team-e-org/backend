CREATE TABLE boards (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    is_private BOOLEAN DEFAULT 0,
    is_archive BOOLEAN DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO
    boards (id, user_id, name, description)
VALUES
    (1, 1, "cat", "cat images"),
    (2, 1, "dog", "dog images"),
    (3, 1, "bird", "bird images");
