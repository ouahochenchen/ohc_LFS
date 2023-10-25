/*
 Navicat Premium Data Transfer

 Source Server         : sk
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : ols_db

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 22/10/2023 16:11:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for lane_order_tab
-- ----------------------------
DROP TABLE IF EXISTS `lane_order_tab`;
CREATE TABLE `lane_order_tab`  (
                                   `order_id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                                   `buyer_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
                                   `buyer_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
                                   `buyer_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
                                   `goods_type` int NOT NULL DEFAULT 0,
                                   `seller_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
                                   `seller_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
                                   `seller_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
                                   `package_height` int NULL DEFAULT 0,
                                   `package_weight` int NULL DEFAULT 0,
                                   `price` decimal(10, 4) NULL DEFAULT NULL,
                                   `order_status` tinyint NOT NULL DEFAULT 0,
                                   `lane_id` int NOT NULL,
                                   PRIMARY KEY (`order_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lane_order_tab
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;








