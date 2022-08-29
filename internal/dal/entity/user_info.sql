CREATE TABLE `user_info` (
  `id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `account` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户账号',
  `password` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户密码',
  `nickname` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户昵称',
  `permission` int NOT NULL DEFAULT '0' COMMENT '用户权限：1、root权限 2、普通权限',
  `create_time` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户信息表'