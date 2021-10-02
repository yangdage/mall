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

 Date: 02/10/2021 21:09:44
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
  `courier_name` char(10) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '物流名称',
  `shipment_number` bigint(20) DEFAULT NULL COMMENT '物流单号',
  `address_id` bigint(20) DEFAULT NULL COMMENT '地址id',
  `user_id` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户id',
  `user_name` char(50) DEFAULT NULL COMMENT '用户名称',
  `admin_id` bigint(20) DEFAULT NULL COMMENT '管理员id',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=202190039031 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order
-- ----------------------------
BEGIN;
INSERT INTO `order` VALUES (202100344553, '103,106,', 8698.80, '待付款', '顺丰', 20213345009, 1002, '202030', 'BBB', 100030, '2021-09-29 16:25:34', NULL);
INSERT INTO `order` VALUES (202100644003, '103,106,', 8698.80, '待付款', '顺丰', 20213345009, 1002, '202030', '李四', 100030, '2021-09-28 18:32:53', '2021-10-01 21:00:26');
INSERT INTO `order` VALUES (202101344528, '107,109,110,', 5748.00, '已发货', '顺丰', 12345, 1002, '202030', 'AAA', 100030, '2021-09-30 18:33:42', '2021-10-01 21:00:48');
INSERT INTO `order` VALUES (202130244550, '107,109,', 5519.22, '已发货', '顺丰', 123, 1002, '202030', '王五', 100030, '2021-09-30 18:33:06', '2021-10-01 20:17:04');
INSERT INTO `order` VALUES (202133405880, '103,', 3308.88, '已完成', '圆通', 20213345009, 1002, '111', 'm100wee', 100030, '2021-10-02 16:16:48', NULL);
INSERT INTO `order` VALUES (202189003242, '103,104,', 3598.34, '已发货', '申通', 12345, 1002, '202030', '张三', 100030, '2021-10-02 16:16:48', '2021-10-01 20:15:17');
INSERT INTO `order` VALUES (202190039023, '103,106,104', 4560.00, '已完成', '圆通', 20213345009, 1002, '123', 'm230ax', 100030, '2021-10-02 16:16:48', NULL);
INSERT INTO `order` VALUES (202190039024, '103,106,1', 3598.34, '待收货', '圆通', 20213345009, 1002, NULL, '张三', 100030, '2021-10-01 16:16:48', NULL);
INSERT INTO `order` VALUES (202190039025, '103,106,104', 3598.34, '待收货', '申通', 20213345009, 1002, NULL, '张三', 100030, '2021-10-02 16:16:48', NULL);
INSERT INTO `order` VALUES (202190039026, '103,106,104', 3598.34, '待评价', '顺丰', 20213345009, 1002, NULL, '张三', 100030, '2021-10-01 16:16:48', NULL);
INSERT INTO `order` VALUES (202190039027, '103,106,', 3598.34, '待评价', '顺丰', 20213345009, 1002, NULL, '张三', 100030, '2021-10-01 16:16:48', NULL);
INSERT INTO `order` VALUES (202190039028, '103,106,1', 3598.34, '待发货', '圆通', 20213345009, 1002, NULL, 'm230ax', 100030, '2021-10-01 16:16:48', NULL);
INSERT INTO `order` VALUES (202190039029, '103,', 3598.34, '待发货', '圆通', 20213345009, 1002, NULL, 'm230ax', 100030, '2021-10-03 16:16:48', NULL);
INSERT INTO `order` VALUES (202190039030, '103,106,104', 3598.34, '已完成', '申通', 20213345009, 1002, NULL, 'm230ax', 100030, '2021-10-02 16:16:48', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
