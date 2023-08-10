DROP DATABASE IF EXISTS testdb;
CREATE DATABASE testdb DEFAULT CHARACTER SET utf8mb4;
USE testdb;

CREATE TABLE testuser (
    id int,
    name varchar(10),
    address varchar(10)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='テスト用テーブル';

INSERT INTO testuser
VALUES (1, 'Yamada', 'Tokyo'),
       (2, 'Tanaka', 'Kanagawa'),
       (3, 'Sato', 'Chiba');
