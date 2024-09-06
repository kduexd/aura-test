CREATE TABLE `user`
(
    `uuid`     char(36)                            NOT NULL COMMENT 'user uuid',
    `name`     varchar(30)                         NOT NULL COMMENT 'user name',
    `password` char(60)                            NOT NULL COMMENT 'password',
    `created`  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT 'create time',
    `updated`  timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL COMMENT 'update time',
    PRIMARY KEY (`uuid`)
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_0900_as_cs COMMENT 'user';