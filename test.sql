/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3306
 Source Schema         : studio_monitor

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 18/05/2023 19:01:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for test
-- ----------------------------
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test`  (
  `id` bigint NOT NULL,
  `user_id` bigint NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `amount` decimal(20, 2) NULL DEFAULT NULL,
  `create_at` datetime NULL DEFAULT NULL,
  `status` smallint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of test
-- ----------------------------
INSERT INTO `test` VALUES (2695795661258752, 1, '广东aa有限公司', 550.00, '2018-05-14 00:00:00', 2);
INSERT INTO `test` VALUES (2695800354586624, 1, '广东aa有限公司', 1210.00, '2018-05-14 00:00:00', 5);
INSERT INTO `test` VALUES (2696688667082752, 1, '东莞cc公司', 47.80, '2018-05-15 09:09:15', 5);
INSERT INTO `test` VALUES (2696755529547776, 1, '广东aa有限公司', 13.00, '2018-05-15 00:00:00', 5);
INSERT INTO `test` VALUES (2696763584970752, 1, '东莞cc公司', 16.42, '2018-05-15 10:30:40', 5);
INSERT INTO `test` VALUES (2696767262343168, 1, '广东aa有限公司', 180.00, '2018-05-15 10:32:37', 5);
INSERT INTO `test` VALUES (2696775042531328, 1, '广东aa有限公司', 100.00, '2018-05-15 10:38:35', 5);
INSERT INTO `test` VALUES (2696790954049536, 1, '东cc公司', 11.75, '2018-05-15 10:53:03', 5);
INSERT INTO `test` VALUES (2696792718098432, 1, '广东aa有限公司', 20.00, '2022-05-04 15:04:59', 5);
INSERT INTO `test` VALUES (2696797306683392, 2, '东莞cc公司', 2.60, '2018-05-15 10:59:29', 4);
INSERT INTO `test` VALUES (2696800312573952, 2, '东莞cc公司', 4890.00, '2018-05-15 00:00:00', 5);
INSERT INTO `test` VALUES (2696839197065216, 2, '东莞cc公司', 117.50, '2018-05-15 11:41:15', 5);
INSERT INTO `test` VALUES (2696992436191232, 2, '东莞cc公司', 2.39, '2018-05-15 14:18:16', 5);
INSERT INTO `test` VALUES (2696996734763008, 2, '东莞cc公司', 2.39, '2018-05-15 14:22:38', 5);
INSERT INTO `test` VALUES (2697019514847232, 2, '广东aa有限公司', 300.00, '2018-05-15 14:47:18', 6);

SET FOREIGN_KEY_CHECKS = 1;
