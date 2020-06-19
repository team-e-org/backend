CREATE TABLE pins_tags (
    id INT PRIMARY KEY AUTO_INCREMENT,
    pin_id INT,
    tag_id INT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO
    pins_tags (id, pin_id, tag_id)
VALUES
    (1, 1, 1),
    (2, 2, 2),
    (3, 3, 3);
