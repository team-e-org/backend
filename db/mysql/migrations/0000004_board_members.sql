CREATE TABLE board_members (
    id INT PRIMARY KEY AUTO_INCREMENT,
    board_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO
    board_members (id, board_id, user_id)
VALUES
    (1, 1, 1),
    (2, 2, 1),
    (3, 3, 1);
