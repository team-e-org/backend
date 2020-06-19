CREATE TABLE boards_pins (
    id INT PRIMARY KEY AUTO_INCREMENT,
    board_id INT NOT NULL,
    pin_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO
    boards_pins (id, board_id, pin_id)
VALUES
    (1, 1, 1),
    (2, 1, 2),
    (3, 1, 3);
