create table if not exists `user`
(
    `id`       varchar(50) not null comment '主键',
    `username` varchar(30) not null comment '姓名',
    `password` varchar(50) not null comment '密码',
    primary key (`id`)
) engine = innodb character set = utf8 comment = '用户表'