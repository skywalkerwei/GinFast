# ************************************************************
# Sequel Ace SQL dump
# 版本号： 20021
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: 127.0.0.1 (MySQL 8.0.18)
# 数据库: skr
# 生成时间: 2022-01-06 14:34:13 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 departments
# ------------------------------------------------------------

DROP TABLE IF EXISTS `departments`;

CREATE TABLE `departments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级部门',
  `title` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部门表';

LOCK TABLES `departments` WRITE;
/*!40000 ALTER TABLE `departments` DISABLE KEYS */;

INSERT INTO `departments` (`id`, `parent_id`, `title`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,0,'技术部','2022-01-01 00:00:00','2022-01-01 00:00:00',NULL),
	(2,0,'产品部','2022-01-01 00:00:00','2022-01-01 00:00:00',NULL),
	(3,0,'销售部','2022-01-01 00:00:00','2022-01-01 00:00:00',NULL);

/*!40000 ALTER TABLE `departments` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 evaluation_parts
# ------------------------------------------------------------

DROP TABLE IF EXISTS `evaluation_parts`;

CREATE TABLE `evaluation_parts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `evaluation_id` int(11) NOT NULL COMMENT '问卷id',
  `title` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '块名称',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='调查问卷块表';



# 转储表 evaluation_questions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `evaluation_questions`;

CREATE TABLE `evaluation_questions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `part_id` int(10) NOT NULL COMMENT '块id',
  `title` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '问题名称',
  `score` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '分数',
  `weight` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '权重',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序',
  `content` json NOT NULL COMMENT '绩效考评标准',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='调查问卷块表';



# 转储表 evaluations
# ------------------------------------------------------------

DROP TABLE IF EXISTS `evaluations`;

CREATE TABLE `evaluations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '问卷名称',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='调查问卷表';



# 转储表 position_evaluation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `position_evaluation`;

CREATE TABLE `position_evaluation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `position_id` int(11) NOT NULL COMMENT '职位id',
  `evaluation_id` int(11) NOT NULL COMMENT '问卷id',
  `weight` int(11) NOT NULL COMMENT '权重',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='职位-问卷绑定关系表';



# 转储表 position_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `position_user`;

CREATE TABLE `position_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '用户id',
  `position_id` int(11) NOT NULL COMMENT '职位id',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='员工职位表';

LOCK TABLES `position_user` WRITE;
/*!40000 ALTER TABLE `position_user` DISABLE KEYS */;

INSERT INTO `position_user` (`id`, `uid`, `position_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,1,1,'2022-01-01 00:00:00','2022-01-01 00:00:00',NULL);

/*!40000 ALTER TABLE `position_user` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 positions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `positions`;

CREATE TABLE `positions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `department_id` int(11) NOT NULL COMMENT '部门',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '职位上级',
  `title` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '职位名称',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部门职位表';

LOCK TABLES `positions` WRITE;
/*!40000 ALTER TABLE `positions` DISABLE KEYS */;

INSERT INTO `positions` (`id`, `department_id`, `parent_id`, `title`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,1,0,'技术总监','2022-01-01 00:00:00','2022-01-01 00:00:00',NULL),
	(2,1,1,'技术经理','2022-01-01 00:00:00','2022-01-01 00:00:00',NULL);

/*!40000 ALTER TABLE `positions` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 respondent
# ------------------------------------------------------------

DROP TABLE IF EXISTS `respondent`;

CREATE TABLE `respondent` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `form_uid` int(11) NOT NULL DEFAULT '0' COMMENT '调查人id',
  `evaluation_id` int(11) NOT NULL DEFAULT '0' COMMENT '问卷id',
  `answer_uid` int(11) NOT NULL DEFAULT '0' COMMENT '填写问卷人id',
  `score` decimal(8,2) NOT NULL COMMENT '分数',
  `content` json NOT NULL COMMENT '填写内容',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '完成状态',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='调查';



# 转储表 user_evaluation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_evaluation`;

CREATE TABLE `user_evaluation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '用户id',
  `evaluation_id` int(11) NOT NULL COMMENT '问卷id',
  `score` decimal(8,2) NOT NULL COMMENT '获取重分数',
  `weight` int(11) NOT NULL COMMENT '权重',
  `created_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户-问卷绑定关系表';



# 转储表 users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名字',
  `avatar` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'sex1 男 2女 0未知',
  `login_at` timestamp NOT NULL DEFAULT '2022-01-01 00:00:00' COMMENT '登录时间',
  `login_ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录IP',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态1正常 0禁用',
  `openid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'openid',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `mobile`, `name`, `avatar`, `sex`, `login_at`, `login_ip`, `status`, `openid`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'18627111095','kyle','',0,'2022-01-06 20:52:25','127.0.0.1',1,'','2021-10-12 00:00:00','2022-01-06 20:52:25',NULL);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
