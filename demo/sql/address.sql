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

 Date: 26/10/2021 15:57:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '地址id',
  `user_id` varchar(200) DEFAULT NULL COMMENT '用户id',
  `name` varchar(255) DEFAULT NULL COMMENT '收货人姓名',
  `mobile` char(16) DEFAULT NULL COMMENT '手机号',
  `postal_code` int(6) DEFAULT NULL COMMENT '邮政编码',
  `province` char(30) DEFAULT NULL COMMENT '省',
  `city` char(30) DEFAULT NULL COMMENT '城市',
  `district` char(30) DEFAULT NULL COMMENT '区/县',
  `detailed_address` varchar(200) DEFAULT NULL COMMENT '详细地址',
  `is_default` tinyint(1) DEFAULT NULL COMMENT '1为默认，0为非默认',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1095 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of address
-- ----------------------------
BEGIN;
INSERT INTO `address` VALUES (1049, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '王俊凯', '13390465582', 339911, '天津市', '天津市', '和平区', '抚平小区209栋33号', 2, '2021-10-25 12:51:53', '2021-10-25 19:03:26');
INSERT INTO `address` VALUES (1050, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '王源', '13599005678', 335600, '北京市', '北京市', '东城区', '呜呜小区1004', 2, '2021-10-25 12:52:34', '2021-10-25 15:40:09');
INSERT INTO `address` VALUES (1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '巴菲特', '13560907668', 133648, '辽宁省', '沈阳市', '和平区', '小区居民', 1, '2021-10-25 14:38:15', '2021-10-26 13:29:47');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
