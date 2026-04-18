CREATE DATABASE IF NOT EXISTS socialnetwork;
USE socialnetwork;

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(50)  NOT NULL,
    nick       VARCHAR(50)  NOT NULL UNIQUE,
    email      VARCHAR(50)  NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE = INNODB;

DROP TABLE IF EXISTS followers;
CREATE TABLE followers
(
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    follower_id INT NOT NULL,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,

    PRIMARY KEY (user_id, follower_id),
    CONSTRAINT chk_no_self_follow CHECK (user_id <> follower_id)
) ENGINE = INNODB;