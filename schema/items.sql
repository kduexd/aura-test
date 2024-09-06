CREATE TABLE `items`
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