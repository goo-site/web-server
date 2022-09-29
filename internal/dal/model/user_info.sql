CREATE TABLE `user_info` (
                             `id` bigint unsigned AUTO_INCREMENT COMMENT 'id',
                             `user_id` varchar(50) NOT NULL DEFAULT '0' COMMENT '用户id',
                             `password` varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码',
                             `email` varchar(50) NOT NULL DEFAULT '' COMMENT '用户邮箱',
                             `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '用户昵称',
                             `permission` tinyint NOT NULL DEFAULT '0' COMMENT '用户权限：1、root权限 2、普通权限',
                             `create_time` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
                             `update_time` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE KEY `user_id_key` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表'