CREATE TABLE IF NOT EXISTS `items`
(
    `id`          int AUTO_INCREMENT           NOT NULL,
    `user_id`     char(36)                     NOT NULL COMMENT 'mapping user uuid',
    `name`        varchar(30)                  NOT NULL COMMENT 'item name',
    `category`    varchar(20)                  NOT NULL COMMENT 'item category',
    `create_time` timestamp                    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `update_time` timestamp                    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='items';

CREATE TABLE IF NOT EXISTS `user`
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