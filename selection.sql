/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3306
 Source Schema         : selection

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 24/05/2023 18:30:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, 'g', '2', 'root', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (24, 'g', '3', 'finance', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (149, 'p', 'finance', '/api/v1/test/tta/:idx', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (44, 'p', 'finance', '/api/v1/user/delete/:id', 'delete', '', '', '');
INSERT INTO `casbin_rule` VALUES (148, 'p', 'systemAdmin', '/api/v1/test/tta/:idx', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (41, 'p', 'systemAdmin', '/api/v1/user/assignRole', 'post', '', '', '');
INSERT INTO `casbin_rule` VALUES (42, 'p', 'systemAdmin', '/api/v1/user/delete/:id', 'delete', '', '', '');
INSERT INTO `casbin_rule` VALUES (40, 'p', 'systemAdmin', '/api/v1/user/list', 'get', '', '', '');

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `perm_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '权限名称',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'URL路径',
  `action` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '权限动作：比如get、post、delete等',
  `perm_type` tinyint NOT NULL DEFAULT 1 COMMENT '权限类型：1-菜单、2-按钮',
  `parent_id` bigint NOT NULL DEFAULT 0 COMMENT '父级ID：资源层级关系',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` bigint NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint NOT NULL DEFAULT 0 COMMENT '更新人',
  `deleted` tinyint NOT NULL DEFAULT 1 COMMENT '是否删除：1-否，2-是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
INSERT INTO `sys_permission` VALUES (4, '用户管理', '/api/v1/user/list', 'get', 1, 0, '2023-05-23 16:12:57', '2023-05-23 16:12:57', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (5, '分配角色', '/api/v1/user/assignRole', 'post', 2, 4, '2023-05-23 16:14:45', '2023-05-23 16:14:45', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (6, '删除用户', '/api/v1/user/delete/:id', 'delete', 2, 4, '2023-05-23 16:15:17', '2023-05-23 16:15:17', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (7, '角色管理', '/api/v1/role/list', 'get', 1, 0, '2023-05-23 16:15:50', '2023-05-23 16:15:50', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (8, '添加角色', '/api/v1/role/add', 'post', 2, 7, '2023-05-23 16:16:31', '2023-05-23 16:16:31', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (9, '分配权限', '/api/v1/role/assignPermission', 'post', 2, 7, '2023-05-23 16:16:56', '2023-05-23 16:16:56', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (10, '删除角色', '/api/v1/role/delete/:id', 'delete', 2, 7, '2023-05-23 16:17:20', '2023-05-23 16:17:20', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (11, '权限管理', '/api/v1/permission/list', 'get', 1, 0, '2023-05-23 16:17:43', '2023-05-23 16:17:43', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (12, '添加权限', '/api/v1/permission/add', 'post', 2, 11, '2023-05-23 16:18:50', '2023-05-23 16:18:50', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (13, '测试', '/api/v1/test/tta/:idx', 'GET', 1, 0, '2023-05-24 15:17:42', '2023-05-24 18:22:23', 2, 2, 1);
INSERT INTO `sys_permission` VALUES (14, '修改权限', '/api/v1/permission/update', 'post', 2, 11, '2023-05-24 18:21:36', '2023-05-24 18:21:36', 2, 2, 1);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '角色编号',
  `role_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` bigint NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint NOT NULL DEFAULT 0 COMMENT '更新人',
  `deleted` tinyint NOT NULL DEFAULT 1 COMMENT '是否删除：1-否，2-是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (2, 'root', '超级管理员', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, 0, 1);
INSERT INTO `sys_role` VALUES (3, 'systemAdmin', '系统管理员', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, 0, 1);
INSERT INTO `sys_role` VALUES (9, 'finance', '财务主管', '2023-05-24 15:07:25', '2023-05-24 15:07:25', 2, 2, 1);

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` bigint NOT NULL DEFAULT 0 COMMENT '角色编号',
  `perm_id` bigint NOT NULL DEFAULT 0 COMMENT '权限ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_permission
-- ----------------------------
INSERT INTO `sys_role_permission` VALUES (60, 3, 4);
INSERT INTO `sys_role_permission` VALUES (61, 3, 5);
INSERT INTO `sys_role_permission` VALUES (62, 3, 6);
INSERT INTO `sys_role_permission` VALUES (63, 3, 13);
INSERT INTO `sys_role_permission` VALUES (64, 9, 6);
INSERT INTO `sys_role_permission` VALUES (65, 9, 13);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `real_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 1-启用，2-禁用',
  `role_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '角色编号',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` bigint NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint NOT NULL DEFAULT 0 COMMENT '更新人',
  `deleted` tinyint NOT NULL DEFAULT 1 COMMENT '是否删除：1-否，2-是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (2, 'phh2008@vip.qq.com', 'phh2008@vip.qq.com', 'phh2008@vip.qq.com', '$2a$10$ITxtKZMlLHEqVQU7x5C62OGyDPiduBNGxKBEZRRJ/jkJnFG2.TSi.', 1, 'root', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, 0, 1);
INSERT INTO `sys_user` VALUES (3, '10000@qq.com', '10000@qq.com', '10000@qq.com', '$2a$10$cKUbSKq3jZFGjYiFQ4wpjukcpZL9tSRO5UolVtpDkPDUah8nR6YLa', 1, 'finance', '2023-05-22 18:51:02', '2023-05-24 15:10:52', 0, 0, 1);
INSERT INTO `sys_user` VALUES (4, '10001@qq.com', '10001@qq.com', '10001@qq.com', '$2a$10$sb6dMahwl85887KtX/ATO.Wob0NsR0UouuRjpOQaEs.qPC2LQxB6q', 1, '', '2023-05-24 09:40:45', '2023-05-24 09:40:45', 2, 2, 1);
INSERT INTO `sys_user` VALUES (5, '10002@qq.com', '10002@qq.com', '10002@qq.com', '$2a$10$i0q/k0Qtc79MYWNHKlpCpu2sPaAKQ3cawWwF0ISEQU7C.nhrvVFfO', 1, '', '2023-05-24 09:41:10', '2023-05-24 09:41:10', 2, 2, 1);
INSERT INTO `sys_user` VALUES (6, '10003@qq.com', '10003@qq.com', '10003@qq.com', '$2a$10$cZ8eslmKEHsvhUOthDvACOvQiduSNxasvnSy8g2jEyxNOgJ6DADV2', 1, '', '2023-05-24 15:08:00', '2023-05-24 15:08:00', 2, 2, 1);

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
