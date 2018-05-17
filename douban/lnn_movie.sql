/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 100125
Source Host           : localhost:3306
Source Database       : movie

Target Server Type    : MYSQL
Target Server Version : 100125
File Encoding         : 65001

Date: 2018-05-17 14:58:36
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for lnn_movie
-- ----------------------------
DROP TABLE IF EXISTS `lnn_movie`;
CREATE TABLE `lnn_movie` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(10) unsigned NOT NULL COMMENT 'movie id',
  `rate` float unsigned NOT NULL COMMENT '评分',
  `title` varchar(100) NOT NULL COMMENT '名称',
  `cover` varchar(200) NOT NULL COMMENT '封面图片',
  `create_time` int(10) unsigned NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
