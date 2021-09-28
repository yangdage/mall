/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : mall

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 28/09/2021 18:39:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for statistics
-- ----------------------------
DROP TABLE IF EXISTS `statistics`;
CREATE TABLE `statistics` (
  `id` bigint(10) NOT NULL COMMENT '统计id',
  `goods_count` int(10) DEFAULT NULL COMMENT '商品数',
  `order_count` int(10) DEFAULT NULL COMMENT '订单数',
  `amount` decimal(20,0) DEFAULT NULL COMMENT '交易金额',
  `visitor_count` int(10) DEFAULT NULL COMMENT '访客数',
  `created` char(30) DEFAULT NULL COMMENT '创建时间',
  `updated` char(30) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of statistics
-- ----------------------------
BEGIN;
INSERT INTO `statistics` VALUES (101, 305, 86, 5890, 2030, NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
