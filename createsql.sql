-- ----------------------------
-- 首页总览数据表
-- ----------------------------
DROP TABLE IF EXISTS `home_page_informations`;
CREATE TABLE `home_page_informations`
(
    `total_visitor` int DEFAULT 0 COMMENT '网站总访问人次'
) charset utf8
  collate utf8_general_ci;
INSERT INTO `home_page_informations`
VALUES (1);

-- ----------------------------
-- 外部网站数据表
-- ----------------------------
DROP TABLE IF EXISTS `outside_webs`;
CREATE TABLE `outside_webs`
(
    `name`        varchar(20) DEFAULT NULL COMMENT '网站内容',
    `link`        varchar(50) DEFAULT NULL COMMENT '网址',
    `class`       varchar(15) DEFAULT NULL COMMENT '学科门类',
    `picture_id`  int         DEFAULT 0 COMMENT '图标id',
    `brief_intro` varchar(20) DEFAULT NULL COMMENT '简介'
) charset utf8
  collate utf8_general_ci;

-- ----------------------------
-- 用户信息数据表
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`          varchar(32) DEFAULT NULL COMMENT '用户uid',
    `account`     varchar(15) DEFAULT NULL COMMENT '学号/工号',
    `password`    varchar(20) DEFAULT NULL COMMENT '密码(sha1)',
    `user_name`   varchar(10) DEFAULT NULL COMMENT '用户名',
    `telephone`   varchar(11) DEFAULT NULL COMMENT '电话号码',
    `qq_number` varchar(15) DEFAULT NULL COMMENT 'QQ号',
    `status`      int         DEFAULT 2 COMMENT '用户权限',
    `college`     int         DEFAULT 1 COMMENT '学院',
    `major`       varchar(15) DEFAULT NULL COMMENT '专业',
    `create_time` varchar(21) DEFAULT NULL COMMENT '创建时间',
    `update_time` varchar(21) DEFAULT NULL COMMENT '修改时间'
) charset utf8
  collate utf8_general_ci;