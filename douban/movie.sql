/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 100125
Source Host           : localhost:3306
Source Database       : movie

Target Server Type    : MYSQL
Target Server Version : 100125
File Encoding         : 65001

Date: 2018-05-23 17:20:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for lnn_alias
-- ----------------------------
DROP TABLE IF EXISTS `lnn_alias`;
CREATE TABLE `lnn_alias` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(10) unsigned NOT NULL,
  `name` varchar(20) NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_district
-- ----------------------------
DROP TABLE IF EXISTS `lnn_district`;
CREATE TABLE `lnn_district` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_language
-- ----------------------------
DROP TABLE IF EXISTS `lnn_language`;
CREATE TABLE `lnn_language` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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

-- ----------------------------
-- Table structure for lnn_movie_district
-- ----------------------------
DROP TABLE IF EXISTS `lnn_movie_district`;
CREATE TABLE `lnn_movie_district` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(10) unsigned NOT NULL,
  `district_id` int(10) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_movie_language
-- ----------------------------
DROP TABLE IF EXISTS `lnn_movie_language`;
CREATE TABLE `lnn_movie_language` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(10) unsigned NOT NULL,
  `language_id` int(10) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_movie_performer
-- ----------------------------
DROP TABLE IF EXISTS `lnn_movie_performer`;
CREATE TABLE `lnn_movie_performer` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(10) unsigned NOT NULL,
  `performer_id` int(10) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_movie_type
-- ----------------------------
DROP TABLE IF EXISTS `lnn_movie_type`;
CREATE TABLE `lnn_movie_type` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(10) unsigned NOT NULL,
  `type_id` int(10) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_performer
-- ----------------------------
DROP TABLE IF EXISTS `lnn_performer`;
CREATE TABLE `lnn_performer` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '演员名',
  `type` tinyint(3) unsigned NOT NULL COMMENT '是否是导演 1：是 0：否',
  `create_time` int(10) unsigned NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_release_date
-- ----------------------------
DROP TABLE IF EXISTS `lnn_release_date`;
CREATE TABLE `lnn_release_date` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(10) unsigned NOT NULL,
  `time` int(10) unsigned NOT NULL,
  `remark` varchar(20) NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_runtime
-- ----------------------------
DROP TABLE IF EXISTS `lnn_runtime`;
CREATE TABLE `lnn_runtime` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `district_id` int(10) unsigned NOT NULL,
  `time` smallint(5) unsigned NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_summary
-- ----------------------------
DROP TABLE IF EXISTS `lnn_summary`;
CREATE TABLE `lnn_summary` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(10) unsigned NOT NULL,
  `text` text NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for lnn_type
-- ----------------------------
DROP TABLE IF EXISTS `lnn_type`;
CREATE TABLE `lnn_type` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `create_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
