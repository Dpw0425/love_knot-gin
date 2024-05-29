CREATE TABLE `users`
(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` varchar(20) NOT NULL COMMENT '用户名',
    `password` varchar(20) NOT NULL COMMENT '密码',
    primary key (`id`)
);