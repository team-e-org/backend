CREATE TABLE pins (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    url TEXT,
    user_id INT,
    image_url TEXT NOT NULL,
    is_private BOOLEAN DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO
    pins (id, user_id, title, description, image_url)
VALUES
    (
        1,
        1,
        "pin1",
        "pin1 description",
        "https://2.bp.blogspot.com/-9wlENBcCWIM/XMZ9zUJZMOI/AAAAAAABSlY/pfuusOBVEz4LVyy3-d21JS1MuzvInErbgCLcBGAs/s800/drink_tapioka_tea_schoolboy.png"
    ),
    (
        2,
        1,
        "pin2",
        "pin2 description",
        "https://2.bp.blogspot.com/-9wlENBcCWIM/XMZ9zUJZMOI/AAAAAAABSlY/pfuusOBVEz4LVyy3-d21JS1MuzvInErbgCLcBGAs/s800/drink_tapioka_tea_schoolboy.png"
    ),
    (
        3,
        1,
        "pin3",
        "pin3 description",
        "https://2.bp.blogspot.com/-9wlENBcCWIM/XMZ9zUJZMOI/AAAAAAABSlY/pfuusOBVEz4LVyy3-d21JS1MuzvInErbgCLcBGAs/s800/drink_tapioka_tea_schoolboy.png"
    );
