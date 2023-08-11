DROP DATABASE IF EXISTS testdb;
CREATE DATABASE testdb DEFAULT CHARACTER SET utf8mb4;
USE testdb;


CREATE TABLE `testuser` (
    `id` INT  NOT NULL PRIMARY KEY
    , `name` VARCHAR(10)
    , `address` VARCHAR(10)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='テスト用テーブル';


INSERT INTO `testuser` (`id`, `name`, `address`)
VALUES
    (1, 'Yamada', 'Tokyo')
    , (2, 'Tanaka', 'Kanagawa')
    , (3, 'Sato', 'Chiba')
;


CREATE TABLE `user` (
    `id` VARCHAR(10) NOT NULL PRIMARY KEY
    , `name` VARCHAR(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';


INSERT INTO `user` (`id`, `name`)
VALUES
    ('1', 'Alice')
    , ('2', 'Bob')
    , ('3', 'Carol')
;


CREATE TABLE `todo` (
    `id` VARCHAR(10) NOT NULL PRIMARY KEY
    , `text` VARCHAR(10) NOT NULL
    , `done` BOOLEAN NOT NULL DEFAULT false
    , `user_id` VARCHAR(10)
    , FOREIGN KEY (`user_id`)
        REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';


INSERT INTO `todo` (`id`, `text`, `done`, `user_id`)
VALUES
    ('1', 'first todo', true, '1')
;
