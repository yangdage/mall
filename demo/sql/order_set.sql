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

 Date: 26/10/2021 15:58:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_set
-- ----------------------------
DROP TABLE IF EXISTS `order_set`;
CREATE TABLE `order_set` (
  `id` int(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单设置id',
  `payment_overtime` int(5) DEFAULT NULL COMMENT '订单付款超时',
  `receive_overtime` int(5) DEFAULT NULL COMMENT '订单收货超时',
  `finish_overtime` int(5) DEFAULT NULL COMMENT '订单完成超时',
  `admin_id` bigint(20) DEFAULT NULL COMMENT '管理员id',
  `created` varchar(20) DEFAULT NULL COMMENT '创建时间',
  `updated` varchar(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order_set
-- ----------------------------
BEGIN;
INSERT INTO `order_set` VALUES (100, 12, 12, 14, NULL, NULL, NULL);
INSERT INTO `order_set` VALUES (104, 12, 12, 7, 100030, '2021-10-21 17:54:36', '2021-10-22 09:18:27');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
