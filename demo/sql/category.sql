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

 Date: 02/10/2021 21:09:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '类目id',
  `name` char(50) DEFAULT NULL COMMENT '类目名称',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '父级类目id',
  `level` int(5) DEFAULT NULL COMMENT '类目层级',
  `sort` int(5) DEFAULT NULL COMMENT '类目排序',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1056 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1032, '手机数码', 0, 1, 66, '', '2021-09-27 18:11:39');
INSERT INTO `category` VALUES (1033, '苹果', 1032, 2, 85, '', '2021-09-27 18:11:52');
INSERT INTO `category` VALUES (1034, 'iPhone13', 1033, 3, 59, '', '2021-09-27 18:12:04');
INSERT INTO `category` VALUES (1035, '家用电器', 0, 1, 12, '', '');
INSERT INTO `category` VALUES (1036, '电冰箱', 1035, 2, 47, '', '');
INSERT INTO `category` VALUES (1053, '文化用品', 0, 1, 55, '2021-09-28 09:07:18', '');
INSERT INTO `category` VALUES (1054, 'IT图书', 1053, 2, 33, '2021-09-28 09:07:42', '');
INSERT INTO `category` VALUES (1055, 'Go入门指南', 1054, 3, 90, '2021-09-28 09:08:09', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
