-- ----------------------------
-- 首页总览数据表
-- ----------------------------
DROP TABLE IF EXISTS `home_page_informations`;
CREATE TABLE `home_page_informations`
(
    `total_visitor` int DEFAULT 0 COMMENT '网站总访问人次'
)charset utf8 collate utf8_general_ci;
INSERT INTO `home_page_informations`
VALUES (1);