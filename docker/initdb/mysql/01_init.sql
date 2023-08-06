DROP DATABASE IF EXISTS testdb;
CREATE DATABASE testdb DEFAULT CHARACTER SET utf8mb4;
USE testdb;

CREATE TABLE user (
    id int,
    name varchar(10),
    address varchar(10)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='テスト用テーブル';

INSERT INTO user
VALUES (1, 'Yamada', 'Tokyo'),
       (2, 'Tanaka', 'Kanagawa'),
       (3, 'Sato', 'Chiba');
