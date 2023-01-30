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
    `picture_id`  int DEFAULT 0 COMMENT '图标id',
    `brief_intro` varchar(20) DEFAULT NULL COMMENT '简介'
) charset utf8
  collate utf8_general_ci;