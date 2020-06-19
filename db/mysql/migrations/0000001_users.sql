CREATE TABLE users (
    id int PRIMARY KEY AUTO_INCREMENT,
    name TEXT NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    icon TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO
    users (id, name, email, password, icon)
VALUES
    (
        1,
        "Shinomiya Kaguya",
        "shinomiya@example.com",
        "$2a$10$SOWUFP.hkVI0CrCJyfh5vuf/Gu.SDpv6Y2DYZ/Dbwyr.AKtlAldFe",
        "https://2.bp.blogspot.com/-9wlENBcCWIM/XMZ9zUJZMOI/AAAAAAABSlY/pfuusOBVEz4LVyy3-d21JS1MuzvInErbgCLcBGAs/s800/drink_tapioka_tea_schoolboy.png"
    )
