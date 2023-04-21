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
    `picture_id`  varchar(32) DEFAULT NULL COMMENT '图标id',
    `brief_intro` varchar(100) DEFAULT NULL COMMENT '简介'
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
    `mailbox`     varchar(25) DEFAULT NULL COMMENT '邮箱地址',
    `status`      int         DEFAULT 2 COMMENT '用户权限',
    `college`     int         DEFAULT 1 COMMENT '学院',
    `major`       varchar(15) DEFAULT NULL COMMENT '专业',
    `create_time` varchar(21) DEFAULT NULL COMMENT '创建时间',
    `update_time` varchar(21) DEFAULT NULL COMMENT '修改时间'
) charset utf8
  collate utf8_general_ci;

-- ----------------------------
-- 博文信息数据表
-- ----------------------------
DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs`
(
    `user_id`     varchar(32) DEFAULT NULL COMMENT '用户uid',
    `blog_id`     varchar(32) DEFAULT NULL COMMENT '博文uid',
    `blog_name`   varchar(15) DEFAULT NULL COMMENT '博文名称',
    `brief_intro` varchar(20) DEFAULT NULL COMMENT '简介',
    `state`       int         DEFAULT 0 COMMENT '初始0 仅自己可见1 软删除2',
    `likes_number`       int         DEFAULT 0 COMMENT '点赞数',
    `create_time` varchar(21) DEFAULT NULL COMMENT '创建时间',
    `update_time` varchar(21) DEFAULT NULL COMMENT '修改时间'
) charset utf8
  collate utf8_general_ci;

-- ----------------------------
-- 帖子信息数据表
-- ----------------------------
DROP TABLE IF EXISTS `bars`;
CREATE TABLE `bars`
(
    `id`           varchar(32) DEFAULT NULL COMMENT '帖子id',
    `user_id`      varchar(32) DEFAULT NULL COMMENT '用户id',
    `classify`     varchar(5)  DEFAULT NULL COMMENT '分类',
    `name`         varchar(15) DEFAULT NULL COMMENT '帖子名称',
    `brief`        varchar(30) DEFAULT NULL COMMENT '简介',
    `create_time`  varchar(21) DEFAULT NULL COMMENT '创建时间',
    `floot_number` int         DEFAULT 1 COMMENT '总盖楼数',
    `page_view`    int         DEFAULT 1 COMMENT '浏览量',
    `likes_Number` int         DEFAULT 0 COMMENT '点赞数'
) charset utf8
  collate utf8_general_ci;

-- ----------------------------
-- 楼层信息数据表
-- ----------------------------
DROP TABLE IF EXISTS `bar_floors`;
CREATE TABLE `bar_floors`
(
    `id`             varchar(32) DEFAULT NULL COMMENT '楼层id',
    `bar_id`         varchar(32) DEFAULT NULL COMMENT '所属帖子id',
    `user_id`        varchar(32) DEFAULT NULL COMMENT '用户id',
    `create_time`    varchar(21) DEFAULT NULL COMMENT '创建时间',
    `likes_Number`   int         DEFAULT 0 COMMENT '点赞数',
    `SonFloorNumber` int         DEFAULT 0 COMMENT '楼中楼数'
) charset utf8
  collate utf8_general_ci;

-- ----------------------------
-- 楼中楼信息数据表
-- ----------------------------
DROP TABLE IF EXISTS `son_floors`;
CREATE TABLE `son_floors`
(
    `id`           varchar(32) DEFAULT NULL COMMENT '楼中楼id',
    `bar_id`       varchar(32) DEFAULT NULL COMMENT '所属帖子id',
    `bar_floor_id` varchar(32) DEFAULT NULL COMMENT '所属楼层id',
    `user_id`      varchar(32) DEFAULT NULL COMMENT '用户id',
    `reply_id`     varchar(32) DEFAULT NULL COMMENT '回复楼中楼id',
    `create_time`  varchar(21) DEFAULT NULL COMMENT '创建时间',
    `likes_Number` int         DEFAULT 0 COMMENT '点赞数'
) charset utf8
  collate utf8_general_ci;