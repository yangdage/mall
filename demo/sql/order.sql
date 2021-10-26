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

 Date: 26/10/2021 15:58:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `product_item` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品项',
  `total_price` decimal(20,2) DEFAULT NULL COMMENT '合计',
  `status` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '未发货' COMMENT '订单状态',
  `courier_name` char(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '物流名称',
  `shipment_number` bigint(20) NOT NULL COMMENT '物流单号',
  `address_id` bigint(20) DEFAULT NULL COMMENT '地址id',
  `user_id` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户id',
  `user_name` char(50) DEFAULT NULL COMMENT '用户名称',
  `admin_id` bigint(20) DEFAULT NULL COMMENT '管理员id',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`courier_name`,`shipment_number`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=202190039052 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order
-- ----------------------------
BEGIN;
INSERT INTO `order` VALUES (202190039034, '103', 8860.00, '待发货', '', 0, 0, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 0, '2021-10-26 14:34:44', '');
INSERT INTO `order` VALUES (202190039035, '103', 8860.00, '待发货', '', 0, 0, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 0, '2021-10-26 14:38:26', '');
INSERT INTO `order` VALUES (202190039036, '103', 8860.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 0, '2021-10-26 14:40:57', '');
INSERT INTO `order` VALUES (202190039037, '103', 8860.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:11:04', '');
INSERT INTO `order` VALUES (202190039038, '103', 8860.00, '已发货', '申通', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:11:42', '2021-10-26 15:29:35');
INSERT INTO `order` VALUES (202190039039, '103', 8860.00, '已发货', '圆通', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:12:29', '2021-10-26 15:31:33');
INSERT INTO `order` VALUES (202190039040, '103', 8860.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:14:54', '');
INSERT INTO `order` VALUES (202190039041, '105,109', 4708.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:16:17', '');
INSERT INTO `order` VALUES (202190039042, '103,105', 11248.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:17:57', '');
INSERT INTO `order` VALUES (202190039043, '', 0.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:18:36', '');
INSERT INTO `order` VALUES (202190039044, '', 0.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:18:41', '');
INSERT INTO `order` VALUES (202190039045, '103,104', 10048.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:19:31', '');
INSERT INTO `order` VALUES (202190039046, '', 0.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:19:41', '');
INSERT INTO `order` VALUES (202190039047, '103,104', 10048.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:21:49', '');
INSERT INTO `order` VALUES (202190039048, '105', 2388.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:22:07', '');
INSERT INTO `order` VALUES (202190039049, '103', 8860.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:23:29', '');
INSERT INTO `order` VALUES (202190039050, '103', 8860.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:24:17', '');
INSERT INTO `order` VALUES (202190039051, '103,104', 10048.00, '待发货', '', 0, 1085, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', '时光不等人', 100030, '2021-10-26 15:37:45', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
