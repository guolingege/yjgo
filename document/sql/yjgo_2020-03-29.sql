# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.26)
# Database: yjgo
# Generation Time: 2020-03-29 12:21:10 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table gen_table
# ------------------------------------------------------------

DROP TABLE IF EXISTS `gen_table`;

CREATE TABLE `gen_table` (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) DEFAULT '' COMMENT '表描述',
  `class_name` varchar(100) DEFAULT '' COMMENT '实体类名称',
  `tpl_category` varchar(200) DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `package_name` varchar(100) DEFAULT NULL COMMENT '生成包路径',
  `module_name` varchar(30) DEFAULT NULL COMMENT '生成模块名',
  `business_name` varchar(30) DEFAULT NULL COMMENT '生成业务名',
  `function_name` varchar(50) DEFAULT NULL COMMENT '生成功能名',
  `function_author` varchar(50) DEFAULT NULL COMMENT '生成功能作者',
  `options` varchar(1000) DEFAULT NULL COMMENT '其它生成选项',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='代码生成业务表';

LOCK TABLES `gen_table` WRITE;
/*!40000 ALTER TABLE `gen_table` DISABLE KEYS */;

INSERT INTO `gen_table` (`table_id`, `table_name`, `table_comment`, `class_name`, `tpl_category`, `package_name`, `module_name`, `business_name`, `function_name`, `function_author`, `options`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(53,'t_answer','回复表','answer','crud','yj-app','module','answer','回复','yunjie','','admin','2020-03-12 18:05:41','',NULL,'');

/*!40000 ALTER TABLE `gen_table` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table gen_table_column
# ------------------------------------------------------------

DROP TABLE IF EXISTS `gen_table_column`;

CREATE TABLE `gen_table_column` (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint(20) DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) DEFAULT NULL COMMENT '列类型',
  `go_type` varchar(500) DEFAULT NULL COMMENT 'Go类型',
  `go_field` varchar(200) DEFAULT NULL COMMENT 'Go字段名',
  `html_field` varchar(200) DEFAULT NULL COMMENT 'html字段名',
  `is_pk` char(1) DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) DEFAULT '' COMMENT '字典类型',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`column_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='代码生成业务表字段';

LOCK TABLES `gen_table_column` WRITE;
/*!40000 ALTER TABLE `gen_table_column` DISABLE KEYS */;

INSERT INTO `gen_table_column` (`column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `html_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `create_by`, `create_time`, `update_by`, `update_time`)
VALUES
	(557,53,'id','主键','bigint(20)','int64','Id','id','1','1','0','1','0','0','0','EQ','input','',1,'admin',NULL,'',NULL),
	(558,53,'pid','问题ID','bigint(20)','int64','Pid','pid','0','0','0','1','1','1','1','EQ','input','',2,'admin',NULL,'',NULL),
	(559,53,'atype','回复人类别','tinyint(1)','int','Atype','atype','0','0','0','1','1','1','1','EQ','select','',3,'admin',NULL,'',NULL),
	(560,53,'user_id','回复人ID','bigint(20)','int64','UserId','userId','0','0','0','1','1','1','1','EQ','input','',4,'admin',NULL,'',NULL),
	(561,53,'nick_name','回复人名称','varchar(50)','string','NickName','nickName','0','0','1','1','1','1','1','LIKE','input','',5,'admin',NULL,'',NULL),
	(562,53,'avatar','回复人头像','varchar(50)','string','Avatar','avatar','0','0','0','1','1','1','1','EQ','input','',6,'admin',NULL,'',NULL),
	(563,53,'remark','回复内容','tinytext','string','Remark','remark','0','0','0','1','1','1','0','EQ','textarea','',7,'admin',NULL,'',NULL),
	(564,53,'img1','回复图片1','varchar(100)','string','Img1','img1','0','0','0','1','1','1','1','EQ','input','',8,'admin',NULL,'',NULL),
	(565,53,'img2','回复图片2','varchar(100)','string','Img2','img2','0','0','0','1','1','1','1','EQ','input','',9,'admin',NULL,'',NULL),
	(566,53,'img3','回复图片3','varchar(100)','string','Img3','img3','0','0','0','1','1','1','1','EQ','input','',10,'admin',NULL,'',NULL),
	(567,53,'status','状态','tinyint(1)','int','Status','status','0','0','1','1','1','1','1','EQ','radio','',11,'admin',NULL,'',NULL),
	(568,53,'create_time','创建时间','datetime','Time','CreateTime','createTime','0','0','0','0','0','0','0','EQ','datatime','',12,'admin',NULL,'',NULL),
	(569,53,'update_time','更新时间','datetime','Time','UpdateTime','updateTime','0','0','0','0','0','0','0','EQ','datatime','',13,'admin',NULL,'',NULL);

/*!40000 ALTER TABLE `gen_table_column` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_config
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_config`;

CREATE TABLE `sys_config` (
  `config_id` int(5) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='参数配置表';

LOCK TABLES `sys_config` WRITE;
/*!40000 ALTER TABLE `sys_config` DISABLE KEYS */;

INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(1,'主框架页-默认皮肤样式名称','sys.index.skinName','skin-blue','Y','admin','2018-03-16 11:33:00','','2020-02-12 15:32:15','蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow'),
	(2,'用户管理-账号初始密码','sys.user.initPassword','123456','Y','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','初始化密码 123456'),
	(3,'主框架页-侧边栏主题','sys.index.sideTheme','theme-dark','Y','admin','2018-03-16 11:33:00','','2020-02-05 10:46:28','深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue'),
	(4,'静态资源网盘存储','sys.resource.url','null','Y','admin','2020-02-18 20:10:33','','2020-02-19 10:36:22','public目录下的静态资源存储到OSS/COS等网盘，如果不存储设为null，设置网址即开始');

/*!40000 ALTER TABLE `sys_config` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dept
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dept`;

CREATE TABLE `sys_dept` (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父部门id',
  `ancestors` varchar(50) DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` char(1) DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='部门表';

LOCK TABLES `sys_dept` WRITE;
/*!40000 ALTER TABLE `sys_dept` DISABLE KEYS */;

INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `del_flag`, `create_by`, `create_time`, `update_by`, `update_time`)
VALUES
	(100,0,'0','云捷网络',0,'admin','15888888888','110@qq.com','0','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00'),
	(110,100,'0,100','运维部门',1,'','','','0','0','admin','2019-12-02 17:07:02','admin','2020-03-29 20:05:44');

/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dict_data
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dict_data`;

CREATE TABLE `sys_dict_data` (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(4) DEFAULT '0' COMMENT '字典排序',
  `dict_label` varchar(100) DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='字典数据表';

LOCK TABLES `sys_dict_data` WRITE;
/*!40000 ALTER TABLE `sys_dict_data` DISABLE KEYS */;

INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(1,1,'男','0','sys_user_sex','','','Y','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','性别男'),
	(2,2,'女','1','sys_user_sex','','','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','性别女'),
	(4,1,'显示','0','sys_show_hide','','primary','Y','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','显示菜单'),
	(5,2,'隐藏','1','sys_show_hide','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','隐藏菜单'),
	(6,1,'正常','0','sys_normal_disable','','primary','Y','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','正常状态'),
	(7,2,'停用','1','sys_normal_disable','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','停用状态'),
	(8,1,'正常','0','sys_job_status','','primary','Y','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','正常状态'),
	(9,2,'暂停','1','sys_job_status','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','停用状态'),
	(10,1,'默认','DEFAULT','sys_job_group','','','Y','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','默认分组'),
	(11,2,'系统','SYSTEM','sys_job_group','','','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统分组'),
	(12,1,'是','Y','sys_yes_no','','primary','Y','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统默认是'),
	(13,2,'否','N','sys_yes_no','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统默认否'),
	(14,1,'通知','1','sys_notice_type','','warning','Y','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','通知'),
	(15,2,'公告','2','sys_notice_type','','success','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','公告'),
	(16,1,'正常','0','sys_notice_status','','primary','Y','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','正常状态'),
	(17,2,'关闭','1','sys_notice_status','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','关闭状态'),
	(18,1,'新增','1','sys_oper_type','','info','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','新增操作'),
	(19,2,'修改','2','sys_oper_type','','info','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','修改操作'),
	(20,3,'删除','3','sys_oper_type','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','删除操作'),
	(21,4,'授权','4','sys_oper_type','','primary','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','授权操作'),
	(22,5,'导出','5','sys_oper_type','','warning','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','导出操作'),
	(23,6,'导入','6','sys_oper_type','','warning','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','导入操作'),
	(24,7,'强退','7','sys_oper_type','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','强退操作'),
	(25,8,'生成代码','8','sys_oper_type','','warning','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','生成操作'),
	(26,9,'清空数据','9','sys_oper_type','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','清空操作'),
	(27,1,'成功','0','sys_common_status','','primary','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','正常状态'),
	(28,2,'失败','1','sys_common_status','','danger','N','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','停用状态'),
	(29,0,'免费用户','0','zjuser_type',NULL,'default','Y','0','admin','2019-12-02 16:56:16','admin',NULL,NULL),
	(30,1,'付费用户','1','zjuser_type',NULL,'primary','Y','0','admin','2019-12-02 16:56:40','admin',NULL,NULL),
	(31,0,'微信用户','0','zxuser_type',NULL,'default','Y','0','admin','2019-12-02 17:14:32','admin',NULL,NULL),
	(32,1,'QQ用户','1','zxuser_type',NULL,'primary','N','0','admin','2019-12-02 17:14:55','admin',NULL,NULL),
	(33,2,'抖音用户','2','zxuser_type',NULL,'primary','N','0','admin','2019-12-02 17:15:21','admin',NULL,NULL);

/*!40000 ALTER TABLE `sys_dict_data` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dict_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dict_type`;

CREATE TABLE `sys_dict_type` (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`),
  UNIQUE KEY `dict_type` (`dict_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='字典类型表';

LOCK TABLES `sys_dict_type` WRITE;
/*!40000 ALTER TABLE `sys_dict_type` DISABLE KEYS */;

INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(1,'用户性别','sys_user_sex','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','用户性别列表'),
	(2,'菜单状态','sys_show_hide','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','菜单状态列表'),
	(3,'系统开关','sys_normal_disable','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统开关列表'),
	(4,'任务状态','sys_job_status','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','任务状态列表'),
	(5,'任务分组','sys_job_group','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','任务分组列表'),
	(6,'系统是否','sys_yes_no','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统是否列表'),
	(7,'通知类型','sys_notice_type','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','通知类型列表'),
	(8,'通知状态','sys_notice_status','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','通知状态列表'),
	(9,'操作类型','sys_oper_type','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','操作类型列表'),
	(10,'系统状态','sys_common_status','0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','登录状态列表'),
	(11,'专家用户类别','zjuser_type','0','admin','2019-12-02 16:55:42','admin',NULL,NULL),
	(12,'咨询用户类别','zxuser_type','0','admin','2019-12-02 17:14:07','admin',NULL,NULL),
	(13,'测试','test3dddd','0','admin','2020-02-05 16:23:06','',NULL,'');

/*!40000 ALTER TABLE `sys_dict_type` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_job
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_job`;

CREATE TABLE `sys_job` (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` varchar(20) DEFAULT '1' COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` char(1) DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`,`job_name`,`job_group`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='定时任务调度表';

LOCK TABLES `sys_job` WRITE;
/*!40000 ALTER TABLE `sys_job` DISABLE KEYS */;

INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(10,'test1','','DEFAULT','test1','* * * * * *','1','1','1','admin','2020-02-26 15:30:27','','2020-02-26 16:51:46',''),
	(12,'test2','helloworld|yjgo','DEFAULT','test2','* * * * * *','1','1','1','admin','2020-02-27 10:20:26','',NULL,'');

/*!40000 ALTER TABLE `sys_job` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_job_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_job_log`;

CREATE TABLE `sys_job_log` (
  `job_log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务日志ID',
  `job_name` varchar(64) NOT NULL COMMENT '任务名称',
  `job_group` varchar(64) NOT NULL COMMENT '任务组名',
  `invoke_target` varchar(500) NOT NULL COMMENT '调用目标字符串',
  `job_message` varchar(500) DEFAULT NULL COMMENT '日志信息',
  `status` char(1) DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
  `exception_info` varchar(2000) DEFAULT '' COMMENT '异常信息',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`job_log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='定时任务调度日志表';



# Dump of table sys_logininfor
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_logininfor`;

CREATE TABLE `sys_logininfor` (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` char(1) DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统访问记录';

LOCK TABLES `sys_logininfor` WRITE;
/*!40000 ALTER TABLE `sys_logininfor` DISABLE KEYS */;

INSERT INTO `sys_logininfor` (`info_id`, `login_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`)
VALUES
	(22,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-24 22:19:09'),
	(23,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-24 22:26:54'),
	(24,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-26 08:40:56'),
	(25,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-27 08:47:37'),
	(26,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-27 08:59:04'),
	(27,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-27 09:00:29'),
	(28,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-27 17:00:10'),
	(29,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-27 18:26:56'),
	(30,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-29 10:05:36'),
	(31,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-30 14:51:57'),
	(32,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-01-31 12:38:31'),
	(33,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-01 15:25:36'),
	(34,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-01 15:25:58'),
	(35,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-01 16:24:10'),
	(36,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-01 19:25:02'),
	(37,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-02 10:49:19'),
	(38,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-02 15:33:37'),
	(39,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-03 12:23:41'),
	(40,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-03 12:39:47'),
	(41,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-03 13:25:39'),
	(42,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-04 13:44:06'),
	(43,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-04 19:23:52'),
	(44,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-05 10:44:57'),
	(45,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-05 16:14:00'),
	(46,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-06 14:05:33'),
	(47,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-06 14:17:41'),
	(48,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-06 14:22:26'),
	(49,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-10 13:16:32'),
	(50,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-10 17:17:56'),
	(51,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-11 16:53:45'),
	(52,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-14 17:23:20'),
	(53,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-16 18:48:24'),
	(54,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-17 15:47:25'),
	(55,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-17 15:48:37'),
	(56,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-17 15:55:01'),
	(57,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-17 16:00:12'),
	(58,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 09:17:26'),
	(59,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 09:18:20'),
	(60,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 11:07:14'),
	(61,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 11:07:36'),
	(62,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 11:08:12'),
	(63,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 11:09:08'),
	(64,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 11:17:34'),
	(65,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 11:17:57'),
	(66,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-18 15:41:46'),
	(67,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-19 09:38:32'),
	(68,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-19 09:41:08'),
	(69,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-19 14:21:37'),
	(70,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-19 14:27:47'),
	(71,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-19 20:22:51'),
	(72,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-19 20:23:11'),
	(73,'admin','175.0.212.81','','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-20 08:47:33'),
	(74,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 08:47:54'),
	(75,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-20 09:04:52'),
	(76,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 09:05:02'),
	(77,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-20 09:06:30'),
	(78,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 09:06:39'),
	(79,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-20 09:09:21'),
	(80,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 09:09:34'),
	(81,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-20 09:13:07'),
	(82,'admin','175.0.212.81','','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 09:13:15'),
	(83,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-20 09:17:29'),
	(84,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-20 09:17:40'),
	(85,'admin','175.0.212.81','深圳','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 09:18:46'),
	(86,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','账号或密码不正确','2020-02-20 10:51:36'),
	(87,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 10:51:48'),
	(88,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 15:31:34'),
	(89,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-20 15:34:24'),
	(90,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-26 14:54:49'),
	(91,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-26 15:21:37'),
	(92,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-02-27 11:08:34'),
	(93,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-09 09:07:10'),
	(94,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-09 09:26:35'),
	(95,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-10 20:59:14'),
	(96,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-10 21:49:57'),
	(97,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-10 21:51:15'),
	(98,'admin111','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-11 11:23:47'),
	(99,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-11 11:24:54'),
	(100,'admin111','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-11 11:25:06'),
	(101,'admin111','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-11 11:26:35'),
	(102,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-12 14:54:58'),
	(103,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-12 18:00:55'),
	(104,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-13 08:26:59'),
	(105,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-14 09:46:34'),
	(106,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-14 09:53:19'),
	(107,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-24 21:16:15'),
	(108,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-29 20:01:12'),
	(109,'admin','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','0','登陆成功','2020-03-29 20:15:58');

/*!40000 ALTER TABLE `sys_logininfor` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(50) NOT NULL COMMENT '菜单名称',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父菜单ID',
  `order_num` int(4) DEFAULT '0' COMMENT '显示顺序',
  `url` varchar(200) DEFAULT '#' COMMENT '请求地址',
  `target` varchar(20) DEFAULT '' COMMENT '打开方式（menuItem页签 menuBlank新窗口）',
  `menu_type` char(1) DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` char(1) DEFAULT '0' COMMENT '菜单状态（0显示 1隐藏）',
  `perms` varchar(100) DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) DEFAULT '#' COMMENT '菜单图标',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='菜单权限表';

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;

INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `url`, `target`, `menu_type`, `visible`, `perms`, `icon`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(1,'系统管理',0,1,'#','','M','0','','fa fa-gear','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统管理目录'),
	(2,'系统监控',0,2,'#','','M','0','','fa fa-video-camera','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统监控目录'),
	(3,'系统工具',0,3,'#','','M','0','','fa fa-bars','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统工具目录'),
	(4,'业务管理',0,1,'#','menuItem','M','0',NULL,'fa fa-newspaper-o','admin','2019-12-02 16:39:09','admin',NULL,''),
	(5,'实例演示',0,5,'#','','M','0','','fa fa-desktop','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统工具目录'),
	(6,'表单演示',5,1,'#','','M','0','','','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','表单演示'),
	(7,'表格演示',5,2,'#','','M','0','','','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','表格演示'),
	(8,'弹框演示',5,3,'#','','M','0','','','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','弹框演示'),
	(9,'操作演示',5,4,'#','','M','0','','','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','操作演示'),
	(10,'报表演示',5,5,'#','','M','0','','','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','报表演示'),
	(11,'图标演示',5,6,'#','','M','0','','','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','图标演示'),
	(12,'栅格演示',6,2,'/demo/form/grid','','C','0','demo:grid:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(13,'下拉框',6,3,'/demo/form/select','','C','0','demo:select:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(14,'时间轴',6,4,'/demo/form/timeline','','C','0','demo:timeline:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(15,'基本表单',6,5,'/demo/form/basic','','C','0','demo:basic:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(16,'卡片列表',6,6,'/demo/form/cards','','C','0','demo:cards:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(17,'功能扩展',6,7,'/demo/form/jasny','','C','0','demo:jasny:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(18,'拖动排序',6,8,'/demo/form/sortable','','C','0','demo:sortable:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(19,'选项卡&面板',6,9,'/demo/form/tabs_panels','','C','0','demo:tabs_panels:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(20,'表单校验',6,10,'/demo/form/validate','','C','0','demo:validate:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(21,'表单向导',6,11,'/demo/form/wizard','','C','0','demo:wizard:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(22,'文件上传',6,12,'/demo/form/upload','','C','0','demo:upload:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(23,'日期和时间',6,13,'/demo/form/datetime','','C','0','demo:datetime:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(24,'富文本编辑器',6,14,'/demo/form/summernote','','C','0','demo:summernote:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(25,'左右互选',6,15,'/demo/form/duallistbox','','C','0','demo:duallistbox:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(27,'按钮演示',6,1,'/demo/form/button','menuItem','C','0','demo:button:view','#','admin','2018-03-16 11:33:00','','2020-02-04 08:46:48',''),
	(28,'数据汇总',7,2,'/demo/table/footer','','C','0','demo:footer:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(29,'组合表头',7,3,'/demo/table/groupHeader','','C','0','demo:groupHeader:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(31,'记住翻页',7,5,'/demo/table/remember','','C','0','demo:remember:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(32,'跳转指定页',7,6,'/demo/table/pageGo','','C','0','demo:pageGo:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(33,'查询参数',7,7,'/demo/table/params','','C','0','demo:params:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(35,'点击加载表格',7,9,'/demo/table/button','','C','0','demo:button:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(36,'表格冻结列',7,10,'/demo/table/fixedColumns','','C','0','demo:fixedColumns:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(37,'触发事件',7,11,'/demo/table/event','','C','0','demo:event:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(38,'细节视图',7,12,'/demo/table/detail','','C','0','demo:detail:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(39,'父子视图',7,13,'/demo/table/child','','C','0','demo:child:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(40,'图片预览',7,14,'/demo/table/image','','C','0','demo:image:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(41,'动态增删改查',7,15,'/demo/table/curd','','C','0','demo:curd:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(42,'表格拖曳',7,16,'/demo/table/recorder','','C','0','demo:recorder:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(43,'行内编辑',7,17,'/demo/table/editable','','C','0','demo:editable:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(44,'其它操作',7,18,'/demo/table/other','','C','0','demo:other:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(45,'查询条件',7,1,'/demo/table/search','','C','0','demo:search:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(46,'弹层组件',8,2,'/demo/modal/layer','','C','0','demo:layer:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(47,'弹层表格',8,3,'/demo/modal/table','','C','0','demo:table:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(48,'模态窗口',8,1,'/demo/modal/dialog','','C','0','demo:dialog:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(49,'其他操作',9,2,'/demo/operate/other','','C','0','demo:other:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(50,'表格操作',9,1,'/demo/operate/table','','C','0','demo:table:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(51,'Peity',10,2,'/demo/report/metrics','','C','0','demo:metrics:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(52,'SparkLine',10,3,'/demo/report/peity','','C','0','demo:peity:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(53,'图表组合',10,4,'/demo/report/sparkline','','C','0','demo:sparkline:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(54,'百度Echarts',10,1,'/demo/report/echarts','','C','0','demo:echarts:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(55,'Glyphicons',11,2,'/demo/icon/glyphicons','','C','0','demo:glyphicons:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(56,'Font Awesome',11,1,'/demo/icon/fontawesome','','C','0','demo:fontawesome:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(100,'用户管理',1,1,'/system/user','','C','0','system:user:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','用户管理菜单'),
	(101,'角色管理',1,2,'/system/role','','C','0','system:role:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','角色管理菜单'),
	(102,'菜单管理',1,3,'/system/menu','','C','0','system:menu:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','菜单管理菜单'),
	(103,'部门管理',1,4,'/system/dept','','C','0','system:dept:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','部门管理菜单'),
	(104,'岗位管理',1,5,'/system/post','','C','0','system:post:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','岗位管理菜单'),
	(105,'字典管理',1,6,'/system/dict','','C','0','system:dict:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','字典管理菜单'),
	(106,'参数设置',1,7,'/system/config','','C','0','system:config:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','参数设置菜单'),
	(107,'通知公告',1,8,'/system/notice','','C','0','system:notice:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','通知公告菜单'),
	(109,'在线用户',2,1,'/monitor/online','','C','0','monitor:online:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','在线用户菜单'),
	(110,'定时任务',2,2,'/monitor/job','','C','0','monitor:job:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','定时任务菜单'),
	(112,'服务监控',2,3,'/monitor/server','','C','0','monitor:server:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','服务监控菜单'),
	(113,'表单构建',3,1,'/tool/build','','C','0','tool:build:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','表单构建菜单'),
	(114,'代码生成',3,2,'/tool/gen','','C','0','tool:gen:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','代码生成菜单'),
	(115,'系统接口',3,3,'/tool/swagger','','C','0','tool:swagger:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','系统接口菜单'),
	(500,'操作日志',2,4,'/monitor/operlog','','C','0','monitor:operlog:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','操作日志菜单'),
	(501,'登录日志',2,5,'/monitor/logininfor','','C','0','monitor:logininfor:view','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','登录日志菜单'),
	(1000,'用户查询',100,1,'#','','F','0','system:user:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1001,'用户新增',100,2,'#','','F','0','system:user:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1002,'用户修改',100,3,'#','','F','0','system:user:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1003,'用户删除',100,4,'#','','F','0','system:user:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1004,'用户导出',100,5,'#','','F','0','system:user:export','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1005,'用户导入',100,6,'#','','F','0','system:user:import','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1006,'重置密码',100,7,'#','','F','0','system:user:resetPwd','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1007,'角色查询',101,1,'#','','F','0','system:role:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1008,'角色新增',101,2,'#','','F','0','system:role:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1009,'角色修改',101,3,'#','','F','0','system:role:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1010,'角色删除',101,4,'#','','F','0','system:role:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1011,'角色导出',101,5,'#','','F','0','system:role:export','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1012,'菜单查询',102,1,'#','','F','0','system:menu:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1013,'菜单新增',102,2,'#','','F','0','system:menu:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1014,'菜单修改',102,3,'#','','F','0','system:menu:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1015,'菜单删除',102,4,'#','','F','0','system:menu:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1016,'部门查询',103,1,'#','','F','0','system:dept:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1017,'部门新增',103,2,'#','','F','0','system:dept:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1018,'部门修改',103,3,'#','','F','0','system:dept:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1019,'部门删除',103,4,'#','','F','0','system:dept:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1020,'岗位查询',104,1,'#','','F','0','system:post:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1021,'岗位新增',104,2,'#','','F','0','system:post:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1022,'岗位修改',104,3,'#','','F','0','system:post:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1023,'岗位删除',104,4,'#','','F','0','system:post:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1024,'岗位导出',104,5,'#','','F','0','system:post:export','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1025,'字典查询',105,1,'#','','F','0','system:dict:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1026,'字典新增',105,2,'#','','F','0','system:dict:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1027,'字典修改',105,3,'#','','F','0','system:dict:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1028,'字典删除',105,4,'#','','F','0','system:dict:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1029,'字典导出',105,5,'#','','F','0','system:dict:export','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1030,'参数查询',106,1,'#','','F','0','system:config:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1031,'参数新增',106,2,'#','','F','0','system:config:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1032,'参数修改',106,3,'#','','F','0','system:config:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1033,'参数删除',106,4,'#','','F','0','system:config:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1034,'参数导出',106,5,'#','','F','0','system:config:export','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1035,'公告查询',107,1,'#','','F','0','system:notice:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1036,'公告新增',107,2,'#','','F','0','system:notice:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1037,'公告修改',107,3,'#','','F','0','system:notice:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1038,'公告删除',107,4,'#','','F','0','system:notice:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1039,'操作查询',500,1,'#','','F','0','monitor:operlog:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1040,'操作删除',500,2,'#','','F','0','monitor:operlog:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1041,'详细信息',500,3,'#','','F','0','monitor:operlog:detail','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1042,'日志导出',500,4,'#','','F','0','monitor:operlog:export','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1043,'登录查询',501,1,'#','','F','0','monitor:logininfor:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1044,'登录删除',501,2,'#','','F','0','monitor:logininfor:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1045,'日志导出',501,3,'#','','F','0','monitor:logininfor:export','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1046,'账户解锁',501,4,'#','','F','0','monitor:logininfor:unlock','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1047,'在线查询',109,1,'#','','F','0','monitor:online:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1048,'批量强退',109,2,'#','','F','0','monitor:online:batchForceLogout','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1049,'单条强退',109,3,'#','','F','0','monitor:online:forceLogout','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1050,'任务查询',110,1,'#','','F','0','monitor:job:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1051,'任务新增',110,2,'#','','F','0','monitor:job:add','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1052,'任务修改',110,3,'#','','F','0','monitor:job:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1053,'任务删除',110,4,'#','','F','0','monitor:job:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1054,'状态修改',110,5,'#','','F','0','monitor:job:changeStatus','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1055,'任务详细',110,6,'#','','F','0','monitor:job:detail','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1056,'任务导出',110,7,'#','','F','0','monitor:job:export','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1057,'生成查询',114,1,'#','','F','0','tool:gen:list','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1058,'生成修改',114,2,'#','','F','0','tool:gen:edit','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1059,'生成删除',114,3,'#','','F','0','tool:gen:remove','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1060,'预览代码',114,4,'#','','F','0','tool:gen:preview','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00',''),
	(1061,'生成代码',114,5,'#','','F','0','tool:gen:code','#','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','');

/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_notice
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_notice`;

CREATE TABLE `sys_notice` (
  `notice_id` int(4) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `notice_title` varchar(50) NOT NULL COMMENT '公告标题',
  `notice_type` char(1) NOT NULL COMMENT '公告类型（1通知 2公告）',
  `notice_content` varchar(2000) DEFAULT NULL COMMENT '公告内容',
  `status` char(1) DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通知公告表';



# Dump of table sys_oper_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_oper_log`;

CREATE TABLE `sys_oper_log` (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) DEFAULT '0' COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) DEFAULT '0' COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) DEFAULT '' COMMENT '操作地点',
  `oper_param` varchar(2000) DEFAULT '' COMMENT '请求参数',
  `json_result` varchar(2000) DEFAULT '' COMMENT '返回参数',
  `status` int(1) DEFAULT '0' COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='操作日志记录';

LOCK TABLES `sys_oper_log` WRITE;
/*!40000 ALTER TABLE `sys_oper_log` DISABLE KEYS */;

INSERT INTO `sys_oper_log` (`oper_id`, `title`, `business_type`, `method`, `request_method`, `operator_type`, `oper_name`, `dept_name`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `json_result`, `status`, `error_msg`, `oper_time`)
VALUES
	(10,'清空操作日志',3,'/monitor/operlog/clean','POST',1,'admin','运维部门','/monitor/operlog/clean','[::1]','内网IP','\"all\"','{\"code\":0,\"msg\":\"操作成功\",\"data\":9,\"otype\":3}',0,'','2020-01-27 08:55:31'),
	(11,'菜单树',0,'/system/menu/roleMenuTreeData','GET',1,'admin','运维部门','/system/menu/roleMenuTreeData','[::1]','内网IP','{\"roleId\":0}','{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":0}',1,'','2020-01-31 12:38:38'),
	(12,'菜单树',0,'/system/menu/roleMenuTreeData','GET',1,'admin','运维部门','/system/menu/roleMenuTreeData','[::1]','内网IP','{\"roleId\":0}','{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":0}',1,'','2020-01-31 12:38:51'),
	(13,'菜单树',0,'/system/menu/roleMenuTreeData','GET',1,'admin','运维部门','/system/menu/roleMenuTreeData','[::1]','内网IP','{\"roleId\":0}','{\"code\":500,\"msg\":\"登陆超时\",\"data\":null,\"otype\":0}',1,'','2020-01-31 12:40:32'),
	(14,'菜单树',0,'/system/menu/roleMenuTreeData','GET',1,'admin','运维部门','/system/menu/roleMenuTreeData','[::1]','内网IP','{\"roleId\":0}','{\"code\":500,\"msg\":\"登陆超时\",\"data\":null,\"otype\":0}',1,'','2020-01-31 12:40:45'),
	(15,'菜单树',0,'/system/menu/roleMenuTreeData','GET',1,'admin','运维部门','/system/menu/roleMenuTreeData','[::1]','内网IP','{\"roleId\":0}','{\"code\":500,\"msg\":\"登陆超时\",\"data\":null,\"otype\":0}',1,'','2020-01-31 12:40:59'),
	(16,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 18:36:47'),
	(17,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 18:37:50'),
	(18,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 18:40:27'),
	(19,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 18:40:35'),
	(20,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 18:45:47'),
	(21,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 18:45:56'),
	(22,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 18:55:55'),
	(23,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 19:01:39'),
	(24,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-01 19:01:42'),
	(25,'新增角色',1,'/system/role/add','POST',1,'admin','运维部门','/system/role/add','[::1]','内网IP','{\"RoleName\":\"普通角色1\",\"RoleKey\":\"common1\",\"RoleSort\":\"3\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1086,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1087,1088,1089,1090,1091,1092\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":3,\"otype\":1}',0,'','2020-02-01 19:05:42'),
	(26,'删除角色',3,'/system/role/remove','POST',1,'admin','运维部门','/system/role/remove','[::1]','内网IP','{\"Ids\":\"3\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-01 19:10:00'),
	(27,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":500,\"msg\":\"不允许操作超级管理员角色\",\"data\":null,\"otype\":0}',1,'','2020-02-02 13:04:32'),
	(28,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":500,\"msg\":\"不允许操作超级管理员角色\",\"data\":null,\"otype\":0}',1,'','2020-02-02 13:05:40'),
	(29,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":500,\"msg\":\"保存数据失败\",\"data\":null,\"otype\":0}',1,'','2020-02-02 13:06:11'),
	(30,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":500,\"msg\":\"保存数据失败\",\"data\":null,\"otype\":0}',1,'','2020-02-02 13:20:01'),
	(31,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":500,\"msg\":\"保存数据失败\",\"data\":null,\"otype\":0}',1,'','2020-02-02 13:20:07'),
	(32,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":500,\"msg\":\"保存数据失败\",\"data\":null,\"otype\":0}',1,'','2020-02-02 13:27:44'),
	(33,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":500,\"msg\":\"保存数据失败\",\"data\":null,\"otype\":0}',1,'','2020-02-02 13:27:49'),
	(34,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":500,\"msg\":\"保存数据失败\",\"data\":null,\"otype\":0}',1,'','2020-02-02 13:27:59'),
	(35,'数据权限保存',0,'/system/role/authDataScope','POST',1,'admin','运维部门','/system/role/authDataScope','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"DataScope\":\"2\",\"DeptIds\":\"100,110\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-02-02 13:29:15'),
	(36,'修改角色',1,'/system/role/edit','POST',1,'admin','运维部门','/system/role/edit','[::1]','内网IP','{\"RoleId\":2,\"RoleName\":\"普通角色\",\"RoleKey\":\"common\",\"RoleSort\":\"2\",\"Status\":\"0\",\"Remark\":\"普通角色\",\"MenuIds\":\"1,100,1001,1002,1003,1004,1005,1006,101,1007,1008,1009,1010,1011,102,1012,1013,1014,1015,103,1016,1017,1018,1019,104,1020,1021,1022,1023,1024,105,1025,1026,1027,1028,1029,106,1030,1031,1032,1033,1034,107,1035,1036,1037,1038,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,112,500,1039,1040,1041,1042,501,1043,1044,1045,1046,3,113,114,1057,1058,1059,1060,1061,115,4,5,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,6,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,7,45,46,47,8,48,49,9,50,51,52,53,10,54,55\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-02 13:32:39'),
	(37,'新增用户',1,'/system/user/add','POST',1,'admin','运维部门','/system/user/add','[::1]','内网IP','{\"UserName\":\"test\",\"Phonenumber\":\"18788996255\",\"Email\":\"ddd@163.com\",\"LoginName\":\"admin111\",\"Password\":\"1qaz2wsx\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"\",\"PostIds\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"用户名称长度为5到30位\",\"data\":null,\"otype\":1}',1,'','2020-02-02 16:03:33'),
	(38,'新增用户',1,'/system/user/add','POST',1,'admin','运维部门','/system/user/add','[::1]','内网IP','{\"UserName\":\"test11\",\"Phonenumber\":\"18788996255\",\"Email\":\"ddd@163.com\",\"LoginName\":\"admin111\",\"Password\":\"1qaz2wsx\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"\",\"PostIds\":\"\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":1}',0,'','2020-02-02 16:03:43'),
	(39,'保存分配用户',1,'/system/role/selectAll','POST',1,'admin','运维部门','/system/role/selectAll','[::1]','内网IP','{\"roleId\":2,\"userIds\":\"2\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-02 16:36:58'),
	(40,'取消用户角色授权',3,'/system/role/cancel','POST',1,'admin','运维部门','/system/role/cancel','[::1]','内网IP','{\"roleId\":2,\"userId\":2}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}',0,'','2020-02-02 16:56:42'),
	(41,'保存分配用户',1,'/system/role/selectAll','POST',1,'admin','运维部门','/system/role/selectAll','[::1]','内网IP','{\"roleId\":2,\"userIds\":\"2\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-02 16:56:47'),
	(42,'取消用户角色授权',3,'/system/role/cancelAll','POST',1,'admin','运维部门','/system/role/cancelAll','[::1]','内网IP','{\"roleId\":2,\"userIds\":\"1,2\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}',0,'','2020-02-02 16:56:51'),
	(43,'保存分配用户',1,'/system/role/selectAll','POST',1,'admin','运维部门','/system/role/selectAll','[::1]','内网IP','{\"roleId\":2,\"userIds\":\"1,2\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-02 16:56:55'),
	(44,'导出Excel',0,'/system/role/export','POST',1,'admin','运维部门','/system/role/export','[::1]','内网IP','{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0,\"SortName\":\"\",\"SortOrder\":\"\"}','{\"code\":500,\"msg\":\"Error 1054: Unknown column \'r.role_id\' in \'field list\', SELECT r.role_id,r.role_name,r.role_key,r.role_sort,r.data_scope,r.status FROM `sys_role` WHERE r.del_flag = \'0\'\\n\",\"data\":null,\"otype\":0}',1,'','2020-02-02 16:57:03'),
	(45,'导出Excel',0,'/system/role/export','POST',1,'admin','运维部门','/system/role/export','[::1]','内网IP','{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0,\"SortName\":\"\",\"SortOrder\":\"\"}','{\"code\":0,\"msg\":\"1580633899644102000.xls\",\"data\":null,\"otype\":0}',0,'','2020-02-02 16:58:19'),
	(46,'删除角色',3,'/system/role/remove','POST',1,'admin','运维部门','/system/role/remove','[::1]','内网IP','{\"Ids\":\"2\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-02 16:59:09'),
	(47,'取消用户角色授权',3,'/system/role/cancel','POST',1,'admin','运维部门','/system/role/cancel','[::1]','内网IP','{\"roleId\":1,\"userId\":1}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}',0,'','2020-02-02 16:59:23'),
	(48,'保存分配用户',1,'/system/role/selectAll','POST',1,'admin','运维部门','/system/role/selectAll','[::1]','内网IP','{\"roleId\":1,\"userIds\":\"1,2\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-02 16:59:32'),
	(49,'新增菜单',1,'/system/menu/add','POST',1,'admin','运维部门','/system/menu/add','[::1]','内网IP','{\"ParentId\":0,\"MenuType\":\"M\",\"MenuName\":\"业务管理1\",\"OrderNum\":5,\"Url\":\"\",\"Icon\":\"fa fa-anchor\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1128,\"otype\":1}',0,'','2020-02-04 08:40:44'),
	(50,'删除菜单',2,'/system/menu/remove?id=1128','GET',1,'admin','运维部门','/system/menu/remove?id=1128','[::1]','内网IP','{\"Ids\":\"\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":2}',1,'','2020-02-04 08:42:51'),
	(51,'删除菜单',2,'/system/menu/remove?id=1128','GET',1,'admin','运维部门','/system/menu/remove?id=1128','[::1]','内网IP','{\"Ids\":\"\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":2}',1,'','2020-02-04 08:43:01'),
	(52,'删除菜单',3,'/system/menu/remove?id=1128','GET',1,'admin','运维部门','/system/menu/remove?id=1128','[::1]','内网IP','{\"id\":1128}','{\"code\":0,\"msg\":\"操作成功\",\"data\":true,\"otype\":3}',0,'','2020-02-04 08:46:14'),
	(53,'修改菜单',1,'/system/menu/edit','POST',1,'admin','运维部门','/system/menu/edit','[::1]','内网IP','{\"MenuId\":11,\"ParentId\":5,\"MenuType\":\"C\",\"MenuName\":\"按钮演示11111\",\"OrderNum\":1,\"Url\":\"/demo/form/button\",\"Icon\":\"#\",\"Target\":\"menuItem\",\"Perms\":\"demo:button:view\",\"Visible\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-04 08:46:35'),
	(54,'修改菜单',1,'/system/menu/edit','POST',1,'admin','运维部门','/system/menu/edit','[::1]','内网IP','{\"MenuId\":11,\"ParentId\":5,\"MenuType\":\"C\",\"MenuName\":\"按钮演示\",\"OrderNum\":1,\"Url\":\"/demo/form/button\",\"Icon\":\"#\",\"Target\":\"menuItem\",\"Perms\":\"demo:button:view\",\"Visible\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-04 08:46:48'),
	(55,'新增部门',1,'/system/dept/add','POST',1,'admin','运维部门','/system/dept/add','[::1]','内网IP','{\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":2,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":111,\"otype\":1}',0,'','2020-02-04 16:15:40'),
	(56,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":3,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":1}',1,'','2020-02-04 16:15:49'),
	(57,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":3,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":1}',1,'','2020-02-04 16:18:37'),
	(58,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":3,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":1}',1,'','2020-02-04 16:19:44'),
	(59,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":0,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":1}',1,'','2020-02-04 16:21:26'),
	(60,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":0,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":1}',1,'','2020-02-04 16:21:28'),
	(61,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":3,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-04 16:22:41'),
	(62,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":3,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-04 16:23:31'),
	(63,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门2\",\"OrderNum\":3,\"Leader\":\"曾尚兵1\",\"Phone\":\"18788996252\",\"Email\":\"ddd111@163.com\",\"Status\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-04 16:24:56'),
	(64,'修改部门',1,'/system/dept/edit','POST',1,'admin','运维部门','/system/dept/edit','[::1]','内网IP','{\"DeptId\":111,\"ParentId\":100,\"DeptName\":\"运维部门2\",\"OrderNum\":3,\"Leader\":\"曾尚兵1\",\"Phone\":\"18788996252\",\"Email\":\"ddd111@163.com\",\"Status\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-04 16:33:12'),
	(65,'删除部门',3,'/system/dept/remove?id=111','GET',1,'admin','运维部门','/system/dept/remove?id=111','[::1]','内网IP','{\"id\":111}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-04 16:33:28'),
	(66,'新增部门',1,'/system/dept/add','POST',1,'admin','运维部门','/system/dept/add','[::1]','内网IP','{\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":5,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":112,\"otype\":1}',0,'','2020-02-04 16:33:41'),
	(67,'新增部门',1,'/system/dept/add','POST',1,'admin','运维部门','/system/dept/add','[::1]','内网IP','{\"ParentId\":100,\"DeptName\":\"运维部门1\",\"OrderNum\":2,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":113,\"otype\":1}',0,'','2020-02-04 16:36:57'),
	(68,'修改岗位',1,'/system/post/edit','POST',1,'admin','运维部门','/system/post/edit','[::1]','内网IP','{\"PostId\":1,\"PostName\":\"董事长\",\"PostCode\":\"ceo\",\"PostSort\":1,\"Status\":\"0\",\"Remark\":\"4223434\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-04 19:36:13'),
	(69,'新增岗位',1,'/system/post/add','POST',1,'admin','运维部门','/system/post/add','[::1]','内网IP','{\"PostName\":\"董事长1\",\"PostCode\":\"ceo1\",\"PostSort\":3,\"Status\":\"0\",\"Remark\":\"wqwqe\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":3,\"otype\":1}',0,'','2020-02-04 19:41:31'),
	(70,'删除岗位',3,'/system/post/remove','POST',1,'admin','运维部门','/system/post/remove','[::1]','内网IP','{\"Ids\":\"3\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-04 19:41:36'),
	(71,'导出Excel',0,'/system/post/export','POST',1,'admin','运维部门','/system/post/export','[::1]','内网IP','{\"PostCode\":\"\",\"Status\":\"\",\"PostName\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0}','{\"code\":0,\"msg\":\"1580816498917507000.xls\",\"data\":null,\"otype\":0}',0,'','2020-02-04 19:41:38'),
	(72,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":3,\"ConfigName\":\"主框架页-侧边栏主题11\",\"ConfigKey\":\"sys.index.sideTheme\",\"ConfigValue\":\"theme-dark\",\"ConfigType\":\"Y\",\"Remark\":\"深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-05 10:46:24'),
	(73,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":3,\"ConfigName\":\"主框架页-侧边栏主题\",\"ConfigKey\":\"sys.index.sideTheme\",\"ConfigValue\":\"theme-dark\",\"ConfigType\":\"Y\",\"Remark\":\"深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-05 10:46:28'),
	(74,'新增参数',1,'/system/config/add','POST',1,'admin','运维部门','/system/config/add','[::1]','内网IP','{\"ConfigName\":\"dfsf\",\"ConfigKey\":\"sdfsdfds\",\"ConfigValue\":\"sdfsdfsd\",\"ConfigType\":\"Y\",\"Remark\":\"sdfsdf\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":5,\"otype\":1}',0,'','2020-02-05 10:46:36'),
	(75,'删除参数',3,'/system/config/remove','POST',1,'admin','运维部门','/system/config/remove','[::1]','内网IP','{\"Ids\":\"5\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-05 10:46:39'),
	(76,'导出Excel',0,'/system/config/export','POST',1,'admin','运维部门','/system/config/export','[::1]','内网IP','{\"ConfigName\":\"\",\"ConfigKey\":\"\",\"ConfigType\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0}','{\"code\":0,\"msg\":\"1580870801582669000.xls\",\"data\":null,\"otype\":0}',0,'','2020-02-05 10:46:41'),
	(77,'导出Excel',0,'/system/role/export','POST',1,'admin','运维部门','/system/role/export','[::1]','内网IP','{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0,\"SortName\":\"\",\"SortOrder\":\"\"}','{\"code\":0,\"msg\":\"1580870904125084000.xls\",\"data\":null,\"otype\":0}',0,'','2020-02-05 10:48:24'),
	(78,'新增字典类型',1,'/system/dict/add','POST',1,'admin','运维部门','/system/dict/add','[::1]','内网IP','{\"DictName\":\"测试\",\"DictType\":\"test3dddd\",\"Status\":\"0\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":13,\"otype\":1}',0,'','2020-02-05 16:23:06'),
	(79,'删除字典数据',3,'/system/dict/data/remove','POST',1,'admin','运维部门','/system/dict/data/remove','[::1]','内网IP','{\"Ids\":\"3\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-05 16:24:41'),
	(80,'新增字典数据',1,'/system/dict/data/add','POST',1,'admin','运维部门','/system/dict/data/add','[::1]','内网IP','{\"DictLabel\":\"测试\",\"DictValue\":\"2\",\"DictType\":\"sys_user_sex\",\"DictSort\":3,\"CssClass\":\"\",\"ListClass\":\"primary\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"样式属性不能为空\",\"data\":null,\"otype\":1}',1,'','2020-02-05 16:53:24'),
	(81,'新增字典数据',1,'/system/dict/data/add','POST',1,'admin','运维部门','/system/dict/data/add','[::1]','内网IP','{\"DictLabel\":\"测试\",\"DictValue\":\"2\",\"DictType\":\"sys_user_sex\",\"DictSort\":3,\"CssClass\":\"\",\"ListClass\":\"primary\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":34,\"otype\":1}',0,'','2020-02-05 16:54:45'),
	(82,'删除字典数据',3,'/system/dict/data/remove','POST',1,'admin','运维部门','/system/dict/data/remove','[::1]','内网IP','{\"Ids\":\"34\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-05 16:54:50'),
	(83,'导出Excel',0,'/system/dict/data/export','POST',1,'admin','运维部门','/system/dict/data/export','[::1]','内网IP','{\"DictType\":\"sys_user_sex\",\"DictLabel\":\"\",\"Status\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0}','{\"code\":0,\"msg\":\"1580892894560263000.xls\",\"data\":null,\"otype\":0}',0,'','2020-02-05 16:54:54'),
	(84,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept,sys_dict_data\"}','{\"code\":500,\"msg\":\"sql: no rows in result set\",\"data\":null,\"otype\":1}',1,'','2020-02-07 11:08:42'),
	(85,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept,sys_dict_data\"}','{\"code\":500,\"msg\":\"sql: no rows in result set\",\"data\":null,\"otype\":1}',1,'','2020-02-07 15:43:46'),
	(86,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept,sys_dict_data\"}','{\"code\":500,\"msg\":\"sql: no rows in result set\",\"data\":null,\"otype\":1}',1,'','2020-02-07 15:44:00'),
	(87,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept,sys_dict_data\"}','{\"code\":500,\"msg\":\"sql: no rows in result set\",\"data\":null,\"otype\":1}',1,'','2020-02-07 15:46:30'),
	(88,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept,sys_dict_data\"}','{\"code\":500,\"msg\":\"sql: no rows in result set\",\"data\":null,\"otype\":1}',1,'','2020-02-07 15:48:52'),
	(89,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept,sys_dict_data\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 16:38:48'),
	(90,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"1\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-07 16:44:07'),
	(91,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 16:46:35'),
	(92,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"3\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-07 16:47:06'),
	(93,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 16:47:56'),
	(94,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"4\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-07 16:48:00'),
	(95,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 16:51:34'),
	(96,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 16:51:46'),
	(97,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"5,6\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}',0,'','2020-02-07 16:51:52'),
	(98,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 16:52:28'),
	(99,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"7\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-07 17:01:57'),
	(100,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 17:02:16'),
	(101,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 17:02:40'),
	(102,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 17:04:15'),
	(103,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 17:05:01'),
	(104,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 18:56:52'),
	(105,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"2,3\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}',0,'','2020-02-07 18:57:59'),
	(106,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 18:59:35'),
	(107,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 18:59:39'),
	(108,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"4,5,6,7\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":4,\"otype\":3}',0,'','2020-02-07 19:02:03'),
	(109,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 19:04:35'),
	(110,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"8,9\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}',0,'','2020-02-07 19:05:17'),
	(111,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config,sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-07 19:05:36'),
	(112,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"10,11\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}',0,'','2020-02-11 14:01:23'),
	(113,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"t_zxuser,t_zjuser\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 14:01:31'),
	(114,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"12,13\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}',0,'','2020-02-11 14:03:20'),
	(115,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"t_zxuser,t_zjuser\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 14:03:34'),
	(116,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"14,15\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}',0,'','2020-02-11 14:38:02'),
	(117,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 14:38:21'),
	(118,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"16\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-11 14:43:24'),
	(119,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 14:43:28'),
	(120,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"17\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-11 14:45:14'),
	(121,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 14:45:18'),
	(122,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"18\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-11 15:07:20'),
	(123,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 15:07:24'),
	(124,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"19\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-11 15:07:35'),
	(125,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 15:07:53'),
	(126,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"20\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-11 15:08:23'),
	(127,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 15:09:37'),
	(128,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 15:09:37'),
	(129,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"21,22\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}',0,'','2020-02-11 15:10:15'),
	(130,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 15:14:13'),
	(131,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"25\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-11 15:15:50'),
	(132,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-11 15:16:22'),
	(133,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:15:22'),
	(134,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:15:29'),
	(135,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:16:18'),
	(136,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:17:41'),
	(137,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:22:18'),
	(138,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:22:29'),
	(139,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:27:04'),
	(140,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:27:13'),
	(141,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\",\"Params\":null,\"Columns\":null}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:29:54'),
	(142,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":1,\"ConfigName\":\"主框架页-默认皮肤样式名称\",\"ConfigKey\":\"sys.index.skinName\",\"ConfigValue\":\"skin-blue\",\"ConfigType\":\"Y\",\"Remark\":\"蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-12 15:30:59'),
	(143,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":1,\"ConfigName\":\"主框架页-默认皮肤样式名称\",\"ConfigKey\":\"sys.index.skinName\",\"ConfigValue\":\"skin-blue\",\"ConfigType\":\"Y\",\"Remark\":\"蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-12 15:32:17'),
	(144,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:33:11'),
	(145,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:35:13'),
	(146,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:36:11'),
	(147,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:37:30'),
	(148,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:46:48'),
	(149,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:52:10'),
	(150,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:52:12'),
	(151,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表名称不能为空; 表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:53:15'),
	(152,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"sys_dept\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:54:26'),
	(153,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"TableId\":0,\"TableName\":\"sys_dept\",\"TableComment\":\"\",\"ClassName\":\"\",\"FunctionAuthor\":\"\",\"TplCategory\":\"\",\"PackageName\":\"\",\"ModuleName\":\"\",\"BusinessName\":\"\",\"FunctionName\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"表描述不能为空; 实体类名称不能为空; 作者不能为空; 生成包路径不能为空; 生成模块名不能为空; 生成业务名不能为空; 生成功能名不能为空\",\"data\":null,\"otype\":2}',1,'','2020-02-12 15:55:40'),
	(154,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','null','{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":2}',1,'','2020-02-12 16:21:17'),
	(155,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','null','{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":2}',1,'','2020-02-12 20:29:28'),
	(156,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','null','{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":2}',1,'','2020-02-12 20:34:17'),
	(157,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','null','{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":2}',1,'','2020-02-12 20:37:42'),
	(158,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','null','{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":2}',1,'','2020-02-12 20:40:36'),
	(159,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-13 08:08:32'),
	(160,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":500,\"msg\":\"Error 1054: Unknown column \'class_name\' in \'field list\', INSERT INTO `gen_table_column`(`class_name`,`business_name`,`create_by`,`update_time`,`table_name`,`module_name`,`update_by`,`table_id`,`package_name`,`function_name`,`options`,`create_time`,`table_comment`,`tpl_category`,`function_author`,`remark`) VALUES(\'dept\',\'dept\',\'admin\',\'2020-02-13 08:08:55\',\'sys_dept\',\'yj-app\',\'admin\',26,\'module\',\'部门\',\'\',\'2020-02-11 15:16:10\',\'部门表\',\'crud\',\'yunjie\',\'\') ON DUPLICATE KEY UPDATE `class_name`=VALUES(`class_name`),`business_name`=VALUES(`business_name`),`create_by`=VALUES(`create_by`),`update_time`=VALUES(`update_time`),`table_name`=VALUES(`table_name`),`module_name`=VALUES(`module_name`),`update_by`=VALUES(`update_by`),`table_id`=VALUES(`table_id`),`package_name`=VALUES(`package_name`),`function_name`=VALUES(`function_name`),`options`=VALUES(`options`),`create_time`=VALUES(`create_time`),`table_comment`=VALUES(`table_comment`),`tpl_category`=VALUES(`tpl_category`),`function_author`=VALUES(`function_author`),`remark`=VALUES(`remark`)\\n\",\"data\":null,\"otype\":2}',1,'','2020-02-13 08:11:11'),
	(161,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-13 08:13:46'),
	(162,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-13 08:14:29'),
	(163,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-13 08:16:55'),
	(164,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-13 08:17:17'),
	(165,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-13 08:17:33'),
	(166,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"26\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-13 08:17:43'),
	(167,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-13 08:17:51'),
	(168,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"27\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-14 18:39:18'),
	(169,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-14 18:39:23'),
	(170,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_config\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-15 09:26:23'),
	(171,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-16 11:51:08'),
	(172,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-16 11:51:31'),
	(173,'生成代码',2,'/tool/gen/edit','POST',1,'admin','运维部门','/tool/gen/edit','[::1]','内网IP','{\"tableName\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-16 20:28:59'),
	(174,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"28,29\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}',0,'','2020-02-16 20:43:22'),
	(175,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-16 20:43:29'),
	(176,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"30\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-16 20:43:51'),
	(177,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-16 20:45:17'),
	(178,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"31\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-16 20:54:03'),
	(179,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-16 20:54:07'),
	(180,'生成代码',0,'/tool/gen/genCode?tableId=32','GET',1,'admin','运维部门','/tool/gen/genCode?tableId=32','[::1]','内网IP','{\"tableId\":32}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-02-16 20:54:10'),
	(181,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"32\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-16 20:58:16'),
	(182,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_dept\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-16 20:58:20'),
	(187,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_user_online\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-17 14:03:51'),
	(207,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"sys_job,sys_job_log\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-18 15:44:13'),
	(208,'删除操作日志',3,'/monitor/operlog/remove','POST',1,'admin','运维部门','/monitor/operlog/remove','[::1]','内网IP','{\"Ids\":\"206,205,204,203,202,201,200,199,198\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":9,\"otype\":3}',0,'','2020-02-18 18:16:19'),
	(209,'删除操作日志',3,'/monitor/operlog/remove','POST',1,'admin','运维部门','/monitor/operlog/remove','[::1]','内网IP','{\"Ids\":\"197,196,195,194,193,192,191,190\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":8,\"otype\":3}',0,'','2020-02-18 18:16:25'),
	(210,'删除操作日志',3,'/monitor/operlog/remove','POST',1,'admin','运维部门','/monitor/operlog/remove','[::1]','内网IP','{\"Ids\":\"189,188,186,185,184,183\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":6,\"otype\":3}',0,'','2020-02-18 18:16:40'),
	(211,'新增参数',1,'/system/config/add','POST',1,'admin','运维部门','/system/config/add','[::1]','内网IP','{\"ConfigName\":\"静态资源OSS存储\",\"ConfigKey\":\"sys.resource.oss\",\"ConfigValue\":\"null\",\"ConfigType\":\"Y\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":4,\"otype\":1}',0,'','2020-02-18 20:10:33'),
	(212,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":4,\"ConfigName\":\"静态资源OSS存储\",\"ConfigKey\":\"sys.resource.oss\",\"ConfigValue\":\"null\",\"ConfigType\":\"Y\",\"Remark\":\"public目录下的静态资源存储到OSS，如果不存储设为null\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-18 20:11:45'),
	(213,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":4,\"ConfigName\":\"静态资源OSS存储\",\"ConfigKey\":\"sys.resource.oss\",\"ConfigValue\":\"http://cos.yunjie.info/yjg/\",\"ConfigType\":\"Y\",\"Remark\":\"public目录下的静态资源存储到OSS，如果不存储设为null\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-18 20:28:33'),
	(214,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":4,\"ConfigName\":\"静态资源网盘存储\",\"ConfigKey\":\"sys.resource.url\",\"ConfigValue\":\"http://cos.yunjie.info/yjg/\",\"ConfigType\":\"Y\",\"Remark\":\"public目录下的静态资源存储到OSS/COS等网盘，如果不存储设为null，设置网址即开始\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-18 20:39:04'),
	(215,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":4,\"ConfigName\":\"静态资源网盘存储\",\"ConfigKey\":\"sys.resource.url\",\"ConfigValue\":\"http://cos.yunjie.info/yjg\",\"ConfigType\":\"Y\",\"Remark\":\"public目录下的静态资源存储到OSS/COS等网盘，如果不存储设为null，设置网址即开始\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-18 20:48:34'),
	(216,'用户强退',0,'/monitor/online/forceLogout','POST',1,'admin','运维部门','/monitor/online/forceLogout','[::1]','内网IP','{\"sessionId\":\"C0LS01Q354HSCZPLNN\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-02-18 21:18:28'),
	(217,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"JobId\":0,\"InvokeTarget\":\"test1()\",\"CronExpression\":\"0 30 * * * *\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"0\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-19 09:34:53'),
	(218,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":4,\"ConfigName\":\"静态资源网盘存储\",\"ConfigKey\":\"sys.resource.url\",\"ConfigValue\":\"null\",\"ConfigType\":\"Y\",\"Remark\":\"public目录下的静态资源存储到OSS/COS等网盘，如果不存储设为null，设置网址即开始http://cos.yunjie.info/yjg\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-19 09:38:03'),
	(219,'修改参数',1,'/system/config/edit','POST',1,'admin','运维部门','/system/config/edit','[::1]','内网IP','{\"ConfigId\":4,\"ConfigName\":\"静态资源网盘存储\",\"ConfigKey\":\"sys.resource.url\",\"ConfigValue\":\"null\",\"ConfigType\":\"Y\",\"Remark\":\"public目录下的静态资源存储到OSS/COS等网盘，如果不存储设为null，设置网址即开始http://cos.yunjie.info/yjg\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-19 10:36:22'),
	(220,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"JobId\":0,\"InvokeTarget\":\"test1()\",\"CronExpression\":\"0 30 * * * *\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-19 11:03:20'),
	(221,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"JobId\":0,\"InvokeTarget\":\"test1()\",\"CronExpression\":\"0 30 * * * *\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-19 11:03:57'),
	(222,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"test\",\"JobGroup\":\"\",\"InvokeTarget\":\"test4\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"状态（0正常 1暂停）不能为空\",\"data\":null,\"otype\":1}',1,'','2020-02-19 11:04:22'),
	(223,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"test\",\"JobGroup\":\"\",\"InvokeTarget\":\"test4\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":4,\"otype\":1}',0,'','2020-02-19 11:04:30'),
	(224,'删除定时任务调度',3,'/monitor/job/remove','POST',1,'admin','运维部门','/monitor/job/remove','[::1]','内网IP','{\"Ids\":\"4\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-19 11:04:42'),
	(225,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"test\",\"JobGroup\":\"\",\"InvokeTarget\":\"testtt\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":5,\"otype\":1}',0,'','2020-02-19 11:06:01'),
	(226,'导出Excel',0,'/monitor/job/export','POST',1,'admin','运维部门','/monitor/job/export','[::1]','内网IP','{\"JobId\":0,\"JobName\":\"\",\"JobGroup\":\"\",\"InvokeTarget\":\"\",\"CronExpression\":\"\",\"MisfirePolicy\":\"\",\"Concurrent\":\"\",\"Status\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0}','{\"code\":0,\"msg\":\"1582081836007320000.xls\",\"data\":null,\"otype\":0}',0,'','2020-02-19 11:10:36'),
	(227,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"JobId\":0,\"InvokeTarget\":\"test1()\",\"CronExpression\":\"0 30 * * * *\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-19 14:04:59'),
	(228,'删除定时任务调度',3,'/monitor/job/remove','POST',1,'admin','运维部门','/monitor/job/remove','[::1]','内网IP','{\"Ids\":\"5\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-19 15:00:40'),
	(229,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"InvokeTarget\":\"sdfsdfsf\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":6,\"otype\":1}',0,'','2020-02-19 15:09:26'),
	(230,'删除定时任务调度',3,'/monitor/job/remove','POST',1,'admin','运维部门','/monitor/job/remove','[::1]','内网IP','{\"Ids\":\"6\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-19 15:09:33'),
	(231,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"InvokeTarget\":\"test1\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":7,\"otype\":1}',0,'','2020-02-19 15:09:44'),
	(232,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"JobId\":0,\"InvokeTarget\":\"test1\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"ddd\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-19 15:09:54'),
	(233,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"JobId\":0,\"InvokeTarget\":\"test1\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"ddd\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-19 15:10:01'),
	(234,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"JobId\":0,\"InvokeTarget\":\"test1\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"ddd\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-02-19 15:11:35'),
	(235,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"JobId\":7,\"InvokeTarget\":\"test1\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"3\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-19 15:12:37'),
	(236,'删除定时任务调度',3,'/monitor/job/remove','POST',1,'admin','运维部门','/monitor/job/remove','[::1]','内网IP','{\"Ids\":\"7\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-19 15:12:42'),
	(237,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"系统默认（无参）\",\"JobGroup\":\"DEFAULT\",\"InvokeTarget\":\"test1\",\"CronExpression\":\"0 0/7 * * * ?\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":8,\"otype\":1}',0,'','2020-02-19 16:08:37'),
	(238,'删除用户',3,'/system/user/remove','POST',1,'admin','运维部门','/system/user/remove','175.0.212.81','深圳','{\"Ids\":\"2\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-20 09:07:05'),
	(239,'修改菜单',1,'/system/menu/edit','POST',1,'admin','运维部门','/system/menu/edit','175.0.212.81','深圳','{\"MenuId\":107,\"ParentId\":1086,\"MenuType\":\"C\",\"MenuName\":\"通知公告\",\"OrderNum\":8,\"Url\":\"/system/notice\",\"Icon\":\"#\",\"Target\":\"menuItem\",\"Perms\":\"system:notice:view\",\"Visible\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-20 09:12:54'),
	(240,'删除用户',3,'/system/user/remove','POST',1,'admin','运维部门','/system/user/remove','175.0.212.81','深圳','{\"Ids\":\"1\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-20 09:17:16'),
	(241,'修改用户密码',2,'/system/user/profile/resetSavePwd','POST',1,'admin','运维部门','/system/user/profile/resetSavePwd','175.0.212.81','','{\"Password\":\"\",\"NewPassword\":\"yunjiego\",\"ConfirmPassword\":\"\"}','{\"code\":500,\"msg\":\"请输入旧密码; 旧密码长度为5到30位; 请输入确认密码; 确认密码长度为5到30位\",\"data\":null,\"otype\":2}',1,'','2020-02-20 09:22:21'),
	(242,'修改用户密码',2,'/system/user/profile/resetSavePwd','POST',1,'admin','运维部门','/system/user/profile/resetSavePwd','175.0.212.81','深圳','{\"Password\":\"\",\"NewPassword\":\"yjg123456\",\"ConfirmPassword\":\"\"}','{\"code\":500,\"msg\":\"请输入旧密码; 旧密码长度为5到30位; 请输入确认密码; 确认密码长度为5到30位\",\"data\":null,\"otype\":2}',1,'','2020-02-20 09:22:54'),
	(243,'修改用户密码',2,'/system/user/profile/resetSavePwd','POST',1,'admin','运维部门','/system/user/profile/resetSavePwd','175.0.212.81','深圳','{\"Password\":\"\",\"NewPassword\":\"yjg123456\",\"ConfirmPassword\":\"\"}','{\"code\":500,\"msg\":\"请输入旧密码; 旧密码长度为5到30位; 请输入确认密码; 确认密码长度为5到30位\",\"data\":null,\"otype\":2}',1,'','2020-02-20 09:23:04'),
	(244,'修改用户密码',2,'/system/user/profile/resetSavePwd','POST',1,'admin','运维部门','/system/user/profile/resetSavePwd','175.0.212.81','深圳','{\"Password\":\"\",\"NewPassword\":\"yjg123456\",\"ConfirmPassword\":\"\"}','{\"code\":500,\"msg\":\"请输入旧密码; 旧密码长度为5到30位; 请输入确认密码; 确认密码长度为5到30位\",\"data\":null,\"otype\":2}',1,'','2020-02-20 09:23:15'),
	(245,'修改用户密码',2,'/system/user/profile/resetSavePwd','POST',1,'admin','运维部门','/system/user/profile/resetSavePwd','[::1]','内网IP','{\"OldPassword\":\"1qaz2wsx\",\"NewPassword\":\"yjg123456\",\"Confirm\":\"yjg123456\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-02-20 09:42:04'),
	(246,'demo演示',2,'/demo/operate/edit','POST',1,'admin','运维部门','/demo/operate/edit','[::1]','内网IP','{\"UserId\":1}','{\"code\":0,\"msg\":\"操作成功\",\"data\":{\"userId\":1,\"userCode\":\"1000001\",\"userPhone\":\"15888888888\",\"userEmail\":\"111@qq.com\",\"userBalance\":100,\"status\":\"0\",\"createTime\":\"2020-01-12 02:02:02\"},\"otype\":2}',0,'','2020-02-21 16:59:28'),
	(247,'demo演示',2,'/demo/operate/edit','POST',1,'admin','运维部门','/demo/operate/edit','[::1]','内网IP','{\"UserId\":1}','{\"code\":0,\"msg\":\"操作成功\",\"data\":{\"userId\":1,\"userCode\":\"1000001\",\"userPhone\":\"15888888888\",\"userEmail\":\"111@qq.com\",\"userBalance\":100,\"status\":\"0\",\"createTime\":\"2020-01-12 02:02:02\"},\"otype\":2}',0,'','2020-02-21 16:59:30'),
	(248,'demo演示',2,'/demo/operate/edit?id=1','GET',1,'admin','运维部门','/demo/operate/edit?id=1','[::1]','内网IP','{\"UserId\":1}','{\"code\":0,\"msg\":\"操作成功\",\"data\":{\"userId\":1,\"userCode\":\"1000001\",\"userPhone\":\"15888888888\",\"userEmail\":\"111@qq.com\",\"userBalance\":100,\"status\":\"0\",\"createTime\":\"2020-01-12 02:02:02\"},\"otype\":2}',0,'','2020-02-21 17:23:28'),
	(249,'demo演示',2,'/demo/operate/edit?id=2','GET',1,'admin','运维部门','/demo/operate/edit?id=2','[::1]','内网IP','{\"UserId\":1}','{\"code\":0,\"msg\":\"操作成功\",\"data\":{\"userId\":1,\"userCode\":\"1000001\",\"userPhone\":\"15888888888\",\"userEmail\":\"111@qq.com\",\"userBalance\":100,\"status\":\"0\",\"createTime\":\"2020-01-12 02:02:02\"},\"otype\":2}',0,'','2020-02-21 17:23:32'),
	(250,'修改角色',1,'/system/role/edit','POST',1,'admin','运维部门','/system/role/edit','[::1]','内网IP','{\"RoleId\":1,\"RoleName\":\"管理员\",\"RoleKey\":\"admin\",\"RoleSort\":\"1\",\"Status\":\"0\",\"Remark\":\"管理员\",\"MenuIds\":\"1,100,502,1001,1004,1005,1098,1099,1101,1103,1104,1105,1106,1107,1108,1109,101,1007,1008,1009,1010,1011,1110,1111,1112,1113,1114,1115,1117,1118,1119,1120,1121,1122,102,1097,1012,1013,1123,1124,1125,1126,1127,103,1016,1017,1018,1019,1093,1094,1116,1129,1130,104,1020,1021,1022,1023,1024,1131,1132,1133,1134,105,1025,1026,1029,1137,1138,1139,1140,1143,1144,1148,1149,106,1030,1031,1034,1135,1136,1086,107,1035,1036,1037,1038,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,1151,1152,112,500,1039,1040,1096,1041,1042,501,1043,1044,1045,1046,1095,3,113,114,1057,1058,1059,1060,1061,1145,1146,1147,1150,115,4,5,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,6,27,28,29,30,31,32,33,35,36,37,38,39,40,41,42,43,44,7,45,46,47,8,48,49,9,50,51,52,53,10,54,55\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-26 14:55:29'),
	(251,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"test1\",\"JobGroup\":\"DEFAULT\",\"JobId\":1,\"InvokeTarget\":\"test1()\",\"CronExpression\":\"0 30 * * * *\",\"MisfirePolicy\":\"2\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-26 15:07:17'),
	(252,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"test1\",\"JobGroup\":\"DEFAULT\",\"JobId\":1,\"InvokeTarget\":\"test1()\",\"CronExpression\":\"0 30 * * * *\",\"MisfirePolicy\":\"2\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-26 15:08:02'),
	(253,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":1}','{\"code\":500,\"msg\":\"任务不存在\",\"data\":null,\"otype\":0}',1,'','2020-02-26 15:21:58'),
	(254,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":1}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 15:23:46'),
	(255,'删除定时任务调度',3,'/monitor/job/remove','POST',1,'admin','运维部门','/monitor/job/remove','[::1]','内网IP','{\"Ids\":\"1,2,3,8\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":4,\"otype\":3}',0,'','2020-02-26 15:27:24'),
	(256,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"test1\",\"JobGroup\":\"DEFAULT\",\"InvokeTarget\":\"test1\",\"CronExpression\":\"* * * * * *\",\"MisfirePolicy\":\"2\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":9,\"otype\":1}',0,'','2020-02-26 15:29:06'),
	(257,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":9}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 15:29:25'),
	(258,'删除定时任务调度',3,'/monitor/job/remove','POST',1,'admin','运维部门','/monitor/job/remove','[::1]','内网IP','{\"Ids\":\"9\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-26 15:29:44'),
	(259,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"test1\",\"JobGroup\":\"DEFAULT\",\"InvokeTarget\":\"test1\",\"CronExpression\":\"* * * * * *\",\"MisfirePolicy\":\"2\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":10,\"otype\":1}',0,'','2020-02-26 15:30:27'),
	(260,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 15:30:31'),
	(261,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 16:28:44'),
	(262,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 16:28:52'),
	(263,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 16:31:55'),
	(264,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 16:33:11'),
	(265,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 16:33:19'),
	(266,'',0,'/monitor/job/run','POST',1,'admin','运维部门','/monitor/job/run','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 16:42:25'),
	(267,'修改定时任务调度',1,'/monitor/job/edit','POST',1,'admin','运维部门','/monitor/job/edit','[::1]','内网IP','{\"JobName\":\"test1\",\"JobGroup\":\"DEFAULT\",\"JobId\":10,\"InvokeTarget\":\"test1\",\"CronExpression\":\"* * * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":1}',0,'','2020-02-26 16:51:46'),
	(268,'',0,'/monitor/job/start','POST',1,'admin','运维部门','/monitor/job/start','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-26 16:52:55'),
	(269,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"test2\",\"JobParams\":\"hello|world\",\"JobGroup\":\"DEFAULT\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"* * * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":11,\"otype\":1}',0,'','2020-02-27 10:18:14'),
	(270,'删除定时任务调度',3,'/monitor/job/remove','POST',1,'admin','运维部门','/monitor/job/remove','[::1]','内网IP','{\"Ids\":\"11\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}',0,'','2020-02-27 10:19:53'),
	(271,'新增定时任务调度',1,'/monitor/job/add','POST',1,'admin','运维部门','/monitor/job/add','[::1]','内网IP','{\"JobName\":\"test2\",\"JobParams\":\"helloworld|yjgo\",\"JobGroup\":\"DEFAULT\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"* * * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":12,\"otype\":1}',0,'','2020-02-27 10:20:26'),
	(272,'',0,'/monitor/job/start','POST',1,'admin','运维部门','/monitor/job/start','[::1]','内网IP','{\"jobId\":12}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 10:20:35'),
	(273,'',0,'/monitor/job/start','POST',1,'admin','运维部门','/monitor/job/start','[::1]','内网IP','{\"jobId\":12}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 10:21:33'),
	(274,'',0,'/monitor/job/start','POST',1,'admin','运维部门','/monitor/job/start','[::1]','内网IP','{\"jobId\":12}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 10:27:15'),
	(275,'',0,'/monitor/job/start','POST',1,'admin','运维部门','/monitor/job/start','[::1]','内网IP','{\"jobId\":12}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 10:28:33'),
	(276,'',0,'/monitor/job/stop','POST',1,'admin','运维部门','/monitor/job/stop','[::1]','内网IP','{\"jobId\":12}','{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 10:28:37'),
	(277,'',0,'/monitor/job/start','POST',1,'admin','运维部门','/monitor/job/start','[::1]','内网IP','{\"jobId\":12}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 10:31:30'),
	(278,'',0,'/monitor/job/start','POST',1,'admin','运维部门','/monitor/job/start','[::1]','内网IP','{\"jobId\":12}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 11:11:12'),
	(279,'',0,'/monitor/job/stop','POST',1,'admin','运维部门','/monitor/job/stop','[::1]','内网IP','{\"jobId\":12}','{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 11:11:18'),
	(280,'',0,'/monitor/job/start','POST',1,'admin','运维部门','/monitor/job/start','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"启动成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 11:11:23'),
	(281,'',0,'/monitor/job/stop','POST',1,'admin','运维部门','/monitor/job/stop','[::1]','内网IP','{\"jobId\":10}','{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}',0,'','2020-02-27 11:11:28'),
	(282,'删除菜单',3,'/system/menu/remove?id=1035','GET',1,'admin','运维部门','/system/menu/remove?id=1035','[::1]','内网IP','{\"id\":1035}','{\"code\":0,\"msg\":\"操作成功\",\"data\":true,\"otype\":3}',0,'','2020-02-28 09:38:45'),
	(283,'删除菜单',3,'/system/menu/remove?id=1036','GET',1,'admin','运维部门','/system/menu/remove?id=1036','[::1]','内网IP','{\"id\":1036}','{\"code\":0,\"msg\":\"操作成功\",\"data\":true,\"otype\":3}',0,'','2020-02-28 09:38:48'),
	(284,'删除菜单',3,'/system/menu/remove?id=1037','GET',1,'admin','运维部门','/system/menu/remove?id=1037','[::1]','内网IP','{\"id\":1037}','{\"code\":0,\"msg\":\"操作成功\",\"data\":true,\"otype\":3}',0,'','2020-02-28 09:38:55'),
	(285,'删除菜单',3,'/system/menu/remove?id=1038','GET',1,'admin','运维部门','/system/menu/remove?id=1038','[::1]','内网IP','{\"id\":1038}','{\"code\":0,\"msg\":\"操作成功\",\"data\":true,\"otype\":3}',0,'','2020-02-28 09:38:57'),
	(286,'删除参数',3,'/tool/gen/remove','POST',1,'admin','运维部门','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"33,34,35,36\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":4,\"otype\":3}',0,'','2020-02-28 13:46:18'),
	(287,'导入表结构',1,'/tool/gen/importTable','POST',1,'admin','运维部门','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"t_answer,t_problem,t_user_withdrawals,t_zjuser,t_zxuser\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}',0,'','2020-02-28 13:53:14'),
	(288,'修改用户',2,'/system/user/edit','POST',1,'admin','运维部门','/system/user/edit','[::1]','内网IP','{\"UserId\":1,\"UserName\":\"超级管理员\",\"Phonenumber\":\"15888888881\",\"Email\":\"dd122221111d@163.com\",\"DeptId\":0,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"1\",\"PostIds\":\"1\",\"Remark\":\"管理员111111\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":1,\"otype\":2}',1,'','2020-03-01 11:42:05'),
	(289,'修改用户',2,'/system/user/edit','POST',1,'admin','运维部门','/system/user/edit','[::1]','内网IP','{\"UserId\":1,\"UserName\":\"超级管理员\",\"Phonenumber\":\"15888888881\",\"Email\":\"dd122221111d@163.com\",\"DeptId\":0,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"1\",\"PostIds\":\"1\",\"Remark\":\"管理员111111\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}',0,'','2020-03-01 11:43:05'),
	(290,'新增用户',1,'/system/user/add','POST',1,'admin','','/system/user/add','[::1]','内网IP','{\"UserName\":\"admin\",\"Phonenumber\":\"18788996255\",\"Email\":\"ddd@163.com\",\"LoginName\":\"admin111\",\"Password\":\"yjg123456\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"1\",\"PostIds\":\"\",\"Remark\":\"\"}','{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":1}',1,'','2020-03-11 11:22:10'),
	(291,'新增用户',1,'/system/user/add','POST',1,'admin','','/system/user/add','[::1]','内网IP','{\"UserName\":\"admin\",\"Phonenumber\":\"18788996255\",\"Email\":\"ddd@163.com\",\"LoginName\":\"admin111\",\"Password\":\"yjg123456\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"1\",\"PostIds\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":3,\"otype\":1}',0,'','2020-03-11 11:22:16'),
	(292,'重置密码',2,'/system/user/resetPwd','POST',1,'admin','','/system/user/resetPwd','[::1]','内网IP','{\"UserId\":3,\"Password\":\"yjg123456\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}',0,'','2020-03-11 11:23:31'),
	(293,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"g_game\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 14:55:06'),
	(294,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"g_user\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 14:55:11'),
	(295,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"g_game\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 14:55:39'),
	(296,'生成代码',3,'/tool/gen/remove','POST',1,'admin','','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"42,43,44\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}',0,'','2020-03-12 14:58:43'),
	(297,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"g_game,g_user\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 14:58:47'),
	(298,'生成代码',3,'/tool/gen/remove','POST',1,'admin','','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"45,46\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}',0,'','2020-03-12 15:00:24'),
	(299,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"g_game\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 15:01:15'),
	(300,'生成代码',3,'/tool/gen/remove','POST',1,'admin','','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"47\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}',0,'','2020-03-12 15:04:45'),
	(301,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"g_game\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 15:05:57'),
	(302,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"g_game\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 15:06:11'),
	(303,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"g_game\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 15:07:52'),
	(304,'生成代码',3,'/tool/gen/remove','POST',1,'admin','','/tool/gen/remove','[::1]','内网IP','{\"Ids\":\"37,38,39,40,41,48,49,50\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}',0,'','2020-03-12 18:01:08'),
	(305,'导入表结构',0,'/tool/gen/importTable','POST',1,'admin','','/tool/gen/importTable','[::1]','内网IP','{\"tables\":\"t_answer\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}',0,'','2020-03-12 18:06:02'),
	(306,'字典数据管理',2,'/system/dict/data/edit','POST',1,'admin','','/system/dict/data/edit','[::1]','内网IP','{\"DictCode\":1,\"DictLabel\":\"男\",\"DictValue\":\"0\",\"DictType\":\"\",\"DictSort\":1,\"CssClass\":\"\",\"ListClass\":\"primary\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"[[*{remark}]]\"}','{\"code\":500,\"msg\":\"字典类型不能为空\",\"data\":null,\"otype\":2}',1,'','2020-03-13 08:27:11'),
	(307,'字典数据管理',2,'/system/dict/data/edit','POST',1,'admin','','/system/dict/data/edit','[::1]','内网IP','{\"DictCode\":1,\"DictLabel\":\"男\",\"DictValue\":\"0\",\"DictType\":\"\",\"DictSort\":1,\"CssClass\":\"\",\"ListClass\":\"primary\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"性别男\"}','{\"code\":500,\"msg\":\"字典类型不能为空\",\"data\":null,\"otype\":2}',1,'','2020-03-13 08:28:10'),
	(308,'字典数据管理',2,'/system/dict/data/edit','POST',1,'admin','','/system/dict/data/edit','[::1]','内网IP','{\"DictCode\":1,\"DictLabel\":\"男\",\"DictValue\":\"0\",\"DictType\":\"\",\"DictSort\":1,\"CssClass\":\"\",\"ListClass\":\"primary\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"性别男\"}','{\"code\":500,\"msg\":\"字典类型不能为空\",\"data\":null,\"otype\":2}',1,'','2020-03-13 08:29:16'),
	(309,'字典数据管理',2,'/system/dict/data/edit','POST',1,'admin','','/system/dict/data/edit','[::1]','内网IP','{\"DictCode\":1,\"DictLabel\":\"男\",\"DictValue\":\"0\",\"DictType\":\"\",\"DictSort\":1,\"CssClass\":\"\",\"ListClass\":\"primary\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"性别男\"}','{\"code\":500,\"msg\":\"字典类型不能为空\",\"data\":null,\"otype\":2}',1,'','2020-03-13 08:49:44'),
	(310,'字典数据管理',2,'/system/dict/data/edit','POST',1,'admin','','/system/dict/data/edit','[::1]','内网IP','{\"DictCode\":1,\"DictLabel\":\"男\",\"DictValue\":\"0\",\"DictType\":\"\",\"DictSort\":1,\"CssClass\":\"\",\"ListClass\":\"success\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"性别男\"}','{\"code\":500,\"msg\":\"字典类型不能为空\",\"data\":null,\"otype\":2}',1,'','2020-03-13 08:51:18'),
	(311,'修改用户',2,'/system/user/edit','POST',1,'admin','','/system/user/edit','[::1]','内网IP','{\"UserId\":3,\"UserName\":\"admin\",\"Phonenumber\":\"18788996255\",\"Email\":\"ddd@163.com\",\"DeptId\":0,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"1\",\"PostIds\":\"1\",\"Remark\":\"\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}',0,'','2020-03-29 20:02:24'),
	(312,'角色管理',2,'/system/role/edit','POST',1,'admin','','/system/role/edit','[::1]','内网IP','{\"RoleId\":1,\"RoleName\":\"管理员\",\"RoleKey\":\"admin\",\"RoleSort\":\"1\",\"Status\":\"0\",\"Remark\":\"管理员\",\"MenuIds\":\"1,100,502,1001,1004,1005,1098,1099,1101,1103,1104,1105,1106,1107,1108,1109,101,1007,1008,1009,1010,1011,1110,1111,1112,1113,1114,1115,1117,1118,1119,1120,1121,1122,102,1097,1012,1013,1123,1124,1125,1126,1127,103,1016,1017,1018,1019,1093,1094,1116,1129,1130,104,1020,1021,1022,1023,1024,1131,1132,1133,1134,105,1025,1026,1029,1137,1138,1139,1140,1143,1144,1148,1149,106,1030,1031,1034,1135,1136,1086,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,1151,1152,112,500,1039,1040,1096,1041,1042,501,1043,1044,1045,1046,1095,3,113,114,1057,1058,1059,1060,1061,1145,1146,1147,1150,115,4,5,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,6,27,28,29,30,31,32,33,35,36,37,38,39,40,41,42,43,44,7,45,46,47,8,48,49,9,50,51,52,53,10,54,55\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}',0,'','2020-03-29 20:02:29'),
	(313,'部门管理',2,'/system/dept/edit','POST',1,'admin','','/system/dept/edit','[::1]','内网IP','{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"运维部门\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\"}','{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}',0,'','2020-03-29 20:05:44');

/*!40000 ALTER TABLE `sys_oper_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_post`;

CREATE TABLE `sys_post` (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) NOT NULL COMMENT '岗位名称',
  `post_sort` int(4) NOT NULL COMMENT '显示顺序',
  `status` char(1) NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='岗位信息表';

LOCK TABLES `sys_post` WRITE;
/*!40000 ALTER TABLE `sys_post` DISABLE KEYS */;

INSERT INTO `sys_post` (`post_id`, `post_code`, `post_name`, `post_sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(1,'ceo','董事长',1,'0','admin','2018-03-16 11:33:00','','2020-02-04 19:36:13','4223434'),
	(2,'se','项目经理',2,'0','admin','2018-03-16 11:33:00','admin','2018-03-16 11:33:00','');

/*!40000 ALTER TABLE `sys_post` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) NOT NULL COMMENT '角色权限字符串',
  `role_sort` int(4) NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `status` char(1) NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色信息表';

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;

INSERT INTO `sys_role` (`role_id`, `role_name`, `role_key`, `role_sort`, `data_scope`, `status`, `del_flag`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(1,'管理员','admin',1,'1','0','0','admin','2018-03-16 11:33:00','','2020-03-29 20:02:29','管理员');

/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role_dept
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role_dept`;

CREATE TABLE `sys_role_dept` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色和部门关联表';

LOCK TABLES `sys_role_dept` WRITE;
/*!40000 ALTER TABLE `sys_role_dept` DISABLE KEYS */;

INSERT INTO `sys_role_dept` (`role_id`, `dept_id`)
VALUES
	(2,100),
	(2,110);

/*!40000 ALTER TABLE `sys_role_dept` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role_menu`;

CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色和菜单关联表';

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;

INSERT INTO `sys_role_menu` (`role_id`, `menu_id`)
VALUES
	(1,1),
	(1,2),
	(1,3),
	(1,4),
	(1,5),
	(1,6),
	(1,7),
	(1,8),
	(1,9),
	(1,10),
	(1,11),
	(1,12),
	(1,13),
	(1,14),
	(1,15),
	(1,16),
	(1,17),
	(1,18),
	(1,19),
	(1,20),
	(1,21),
	(1,22),
	(1,23),
	(1,24),
	(1,25),
	(1,26),
	(1,27),
	(1,28),
	(1,29),
	(1,30),
	(1,31),
	(1,32),
	(1,33),
	(1,35),
	(1,36),
	(1,37),
	(1,38),
	(1,39),
	(1,40),
	(1,41),
	(1,42),
	(1,43),
	(1,44),
	(1,45),
	(1,46),
	(1,47),
	(1,48),
	(1,49),
	(1,50),
	(1,51),
	(1,52),
	(1,53),
	(1,54),
	(1,55),
	(1,100),
	(1,101),
	(1,102),
	(1,103),
	(1,104),
	(1,105),
	(1,106),
	(1,109),
	(1,110),
	(1,112),
	(1,113),
	(1,114),
	(1,115),
	(1,500),
	(1,501),
	(1,502),
	(1,1001),
	(1,1004),
	(1,1005),
	(1,1007),
	(1,1008),
	(1,1009),
	(1,1010),
	(1,1011),
	(1,1012),
	(1,1013),
	(1,1016),
	(1,1017),
	(1,1018),
	(1,1019),
	(1,1020),
	(1,1021),
	(1,1022),
	(1,1023),
	(1,1024),
	(1,1025),
	(1,1026),
	(1,1029),
	(1,1030),
	(1,1031),
	(1,1034),
	(1,1039),
	(1,1040),
	(1,1041),
	(1,1042),
	(1,1043),
	(1,1044),
	(1,1045),
	(1,1046),
	(1,1047),
	(1,1048),
	(1,1049),
	(1,1050),
	(1,1051),
	(1,1052),
	(1,1053),
	(1,1054),
	(1,1055),
	(1,1056),
	(1,1057),
	(1,1058),
	(1,1059),
	(1,1060),
	(1,1061),
	(1,1086),
	(1,1093),
	(1,1094),
	(1,1095),
	(1,1096),
	(1,1097),
	(1,1098),
	(1,1099),
	(1,1101),
	(1,1103),
	(1,1104),
	(1,1105),
	(1,1106),
	(1,1107),
	(1,1108),
	(1,1109),
	(1,1110),
	(1,1111),
	(1,1112),
	(1,1113),
	(1,1114),
	(1,1115),
	(1,1116),
	(1,1117),
	(1,1118),
	(1,1119),
	(1,1120),
	(1,1121),
	(1,1122),
	(1,1123),
	(1,1124),
	(1,1125),
	(1,1126),
	(1,1127),
	(1,1129),
	(1,1130),
	(1,1131),
	(1,1132),
	(1,1133),
	(1,1134),
	(1,1135),
	(1,1136),
	(1,1137),
	(1,1138),
	(1,1139),
	(1,1140),
	(1,1143),
	(1,1144),
	(1,1145),
	(1,1146),
	(1,1147),
	(1,1148),
	(1,1149),
	(1,1150),
	(1,1151),
	(1,1152),
	(2,1),
	(2,2),
	(2,3),
	(2,4),
	(2,5),
	(2,6),
	(2,7),
	(2,8),
	(2,9),
	(2,10),
	(2,11),
	(2,12),
	(2,13),
	(2,14),
	(2,15),
	(2,16),
	(2,17),
	(2,18),
	(2,19),
	(2,20),
	(2,21),
	(2,22),
	(2,23),
	(2,24),
	(2,25),
	(2,26),
	(2,27),
	(2,28),
	(2,29),
	(2,30),
	(2,31),
	(2,32),
	(2,33),
	(2,34),
	(2,35),
	(2,36),
	(2,37),
	(2,38),
	(2,39),
	(2,40),
	(2,41),
	(2,42),
	(2,43),
	(2,44),
	(2,45),
	(2,46),
	(2,47),
	(2,48),
	(2,49),
	(2,50),
	(2,51),
	(2,52),
	(2,53),
	(2,54),
	(2,55),
	(2,100),
	(2,101),
	(2,102),
	(2,103),
	(2,104),
	(2,105),
	(2,106),
	(2,107),
	(2,109),
	(2,110),
	(2,112),
	(2,113),
	(2,114),
	(2,115),
	(2,500),
	(2,501),
	(2,1001),
	(2,1002),
	(2,1003),
	(2,1004),
	(2,1005),
	(2,1006),
	(2,1007),
	(2,1008),
	(2,1009),
	(2,1010),
	(2,1011),
	(2,1012),
	(2,1013),
	(2,1014),
	(2,1015),
	(2,1016),
	(2,1017),
	(2,1018),
	(2,1019),
	(2,1020),
	(2,1021),
	(2,1022),
	(2,1023),
	(2,1024),
	(2,1025),
	(2,1026),
	(2,1027),
	(2,1028),
	(2,1029),
	(2,1030),
	(2,1031),
	(2,1032),
	(2,1033),
	(2,1034),
	(2,1035),
	(2,1036),
	(2,1037),
	(2,1038),
	(2,1039),
	(2,1040),
	(2,1041),
	(2,1042),
	(2,1043),
	(2,1044),
	(2,1045),
	(2,1046),
	(2,1047),
	(2,1048),
	(2,1049),
	(2,1050),
	(2,1051),
	(2,1052),
	(2,1053),
	(2,1054),
	(2,1055),
	(2,1056),
	(2,1057),
	(2,1058),
	(2,1059),
	(2,1060),
	(2,1061),
	(3,1062),
	(3,1063),
	(3,1064),
	(3,1065),
	(3,1066),
	(3,1067),
	(3,1068),
	(3,1069),
	(3,1070),
	(3,1071),
	(3,1072),
	(3,1073),
	(3,1074),
	(3,1075),
	(3,1076),
	(3,1077),
	(3,1078),
	(3,1079),
	(3,1080),
	(3,1081),
	(3,1082),
	(3,1083),
	(3,1084),
	(3,1085),
	(3,1086),
	(3,1087),
	(3,1088),
	(3,1089),
	(3,1090),
	(3,1091),
	(3,1092);

/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门ID',
  `login_name` varchar(30) NOT NULL COMMENT '登录账号',
  `user_name` varchar(30) NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) DEFAULT '00' COMMENT '用户类型（00系统用户）',
  `email` varchar(50) DEFAULT '' COMMENT '用户邮箱',
  `phonenumber` varchar(11) DEFAULT '' COMMENT '手机号码',
  `sex` char(1) DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) DEFAULT '' COMMENT '头像路径',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `salt` varchar(20) DEFAULT '' COMMENT '盐加密',
  `status` char(1) DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `login_ip` varchar(50) DEFAULT '' COMMENT '最后登陆IP',
  `login_date` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息表';

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;

INSERT INTO `sys_user` (`user_id`, `dept_id`, `login_name`, `user_name`, `user_type`, `email`, `phonenumber`, `sex`, `avatar`, `password`, `salt`, `status`, `del_flag`, `login_ip`, `login_date`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`)
VALUES
	(1,0,'admin','超级管理员','00','dd122221111d@163.com','15888888881','0','/upload/admin1579180516186761000.png','609ac514e6ef0b9a5f4eee66693fd7f8','NcSB3H','0','0','127.0.0.1','2020-01-13 13:20:40','admin','2018-03-16 11:33:00','admin','2020-03-01 11:43:05','管理员111111'),
	(3,0,'admin111','admin','','ddd@163.com','18788996255','0','','d505d06b9e2645026018b0f1af9cbb8c','EsnR7i','0','0','',NULL,'admin','2020-03-11 11:22:16','admin','2020-03-29 20:02:24','');

/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user_online
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user_online`;

CREATE TABLE `sys_user_online` (
  `sessionId` varchar(50) NOT NULL DEFAULT '' COMMENT '用户会话id',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `dept_name` varchar(50) DEFAULT '' COMMENT '部门名称',
  `ipaddr` varchar(50) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` varchar(10) DEFAULT '' COMMENT '在线状态on_line在线off_line离线',
  `start_timestamp` datetime DEFAULT NULL COMMENT 'session创建时间',
  `last_access_time` datetime DEFAULT NULL COMMENT 'session最后访问时间',
  `expire_time` int(5) DEFAULT '0' COMMENT '超时时间，单位为分钟',
  PRIMARY KEY (`sessionId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='在线用户记录';

LOCK TABLES `sys_user_online` WRITE;
/*!40000 ALTER TABLE `sys_user_online` DISABLE KEYS */;

INSERT INTO `sys_user_online` (`sessionId`, `login_name`, `dept_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `start_timestamp`, `last_access_time`, `expire_time`)
VALUES
	('C1A6G4JTFKI0HAFO20','admin','','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','on_line','2020-03-14 09:46:34','2020-03-14 09:46:34',1440),
	('C1J3DMZYYVOWS7S7WJ','admin','','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','on_line','2020-03-24 21:16:15','2020-03-24 21:16:15',1440),
	('C1NAWWEVNJ4ODV45AN','admin','','[::1]','内网IP','Chrome','Intel Mac OS X 10_14_6','on_line','2020-03-29 20:01:12','2020-03-29 20:01:12',1440);

/*!40000 ALTER TABLE `sys_user_online` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user_post`;

CREATE TABLE `sys_user_post` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户与岗位关联表';

LOCK TABLES `sys_user_post` WRITE;
/*!40000 ALTER TABLE `sys_user_post` DISABLE KEYS */;

INSERT INTO `sys_user_post` (`user_id`, `post_id`)
VALUES
	(1,1),
	(2,2),
	(3,1),
	(8,1),
	(8,2);

/*!40000 ALTER TABLE `sys_user_post` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user_role`;

CREATE TABLE `sys_user_role` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户和角色关联表';

LOCK TABLES `sys_user_role` WRITE;
/*!40000 ALTER TABLE `sys_user_role` DISABLE KEYS */;

INSERT INTO `sys_user_role` (`user_id`, `role_id`)
VALUES
	(1,1),
	(2,1),
	(2,2),
	(3,1);

/*!40000 ALTER TABLE `sys_user_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_answer
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_answer`;

CREATE TABLE `t_answer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` bigint(20) DEFAULT NULL COMMENT '问题ID',
  `atype` tinyint(1) DEFAULT '0' COMMENT '回复人类别',
  `user_id` bigint(20) DEFAULT NULL COMMENT '回复人ID',
  `nick_name` varchar(50) DEFAULT '' COMMENT '回复人名称',
  `avatar` varchar(50) DEFAULT '' COMMENT '回复人头像',
  `remark` tinytext COMMENT '回复内容',
  `img1` varchar(100) DEFAULT '' COMMENT '回复图片1',
  `img2` varchar(100) DEFAULT '' COMMENT '回复图片2',
  `img3` varchar(100) DEFAULT '' COMMENT '回复图片3',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='回复表';



# Dump of table t_problem
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_problem`;

CREATE TABLE `t_problem` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `zx_id` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `zx_name` varchar(30) DEFAULT '' COMMENT '用户昵称',
  `zx_avatar` varchar(50) DEFAULT '' COMMENT '咨询头像',
  `zj_id` bigint(20) DEFAULT NULL COMMENT '专家ID',
  `remark` varchar(250) DEFAULT '' COMMENT '问题描述',
  `img1` varchar(100) DEFAULT '' COMMENT '问题图片1',
  `img2` varchar(100) DEFAULT '' COMMENT '问题图片2',
  `img3` varchar(100) DEFAULT '' COMMENT '问题图片3',
  `ispay` tinyint(1) DEFAULT '0' COMMENT '是否付费',
  `pay_time` datetime DEFAULT NULL COMMENT '付费时间',
  `pay_no` varchar(50) DEFAULT '' COMMENT '订单号',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='问题表';



# Dump of table t_user_withdrawals
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_user_withdrawals`;

CREATE TABLE `t_user_withdrawals` (
  `wid` int(50) NOT NULL AUTO_INCREMENT COMMENT '申请ID',
  `uid` int(50) DEFAULT '0' COMMENT '用户ID',
  `realname` varchar(50) DEFAULT NULL COMMENT '姓名',
  `idno` varchar(50) DEFAULT NULL COMMENT '身份证号',
  `wmoney` decimal(10,0) DEFAULT '0' COMMENT '提现金额',
  `wbankname` varchar(50) DEFAULT '' COMMENT '银行名称',
  `wbankaccount` varchar(50) DEFAULT '' COMMENT '银行帐户',
  `wrealname` varchar(50) DEFAULT '' COMMENT '银行户名',
  `wopenbank` varchar(50) DEFAULT '' COMMENT '开户行',
  `wcreatetime` varchar(50) DEFAULT '' COMMENT '申请时间',
  `wstatus` int(11) DEFAULT '0' COMMENT '审核结果',
  `wremark` varchar(255) DEFAULT '' COMMENT '审核意见',
  `waudittime` varchar(50) DEFAULT '' COMMENT '审核时间',
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='提现申请';



# Dump of table t_zjuser
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_zjuser`;

CREATE TABLE `t_zjuser` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uname` varchar(11) DEFAULT '' COMMENT '手机号',
  `puid` bigint(20) DEFAULT NULL COMMENT '推荐人ID',
  `puname` varchar(11) DEFAULT '' COMMENT '推荐人手机号',
  `real_name` varchar(30) DEFAULT '' COMMENT '姓名',
  `idno` varchar(18) DEFAULT NULL COMMENT '身份证号',
  `avatar` varchar(50) DEFAULT '' COMMENT '头像',
  `job` varchar(50) DEFAULT '' COMMENT '职业',
  `utype` tinyint(1) DEFAULT '0' COMMENT '用户类别',
  `upwd` varchar(50) DEFAULT '' COMMENT '密码',
  `salt` varchar(45) DEFAULT '' COMMENT '密码盐',
  `idpic1` varchar(100) DEFAULT NULL COMMENT '身份证正面',
  `idpic2` varchar(100) DEFAULT NULL COMMENT '身份证反面',
  `pimg1` varchar(100) DEFAULT NULL COMMENT '职业资格认证1',
  `pimg2` varchar(100) DEFAULT NULL COMMENT '职业资格认证2',
  `pimg3` varchar(100) DEFAULT NULL COMMENT '职业资格认证3',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态',
  `im_money` decimal(8,2) DEFAULT '0.00' COMMENT '在线咨询费用',
  `tel_money` decimal(8,2) DEFAULT '0.00' COMMENT '电话咨询费用',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='专家用户';

LOCK TABLES `t_zjuser` WRITE;
/*!40000 ALTER TABLE `t_zjuser` DISABLE KEYS */;

INSERT INTO `t_zjuser` (`id`, `uname`, `puid`, `puname`, `real_name`, `idno`, `avatar`, `job`, `utype`, `upwd`, `salt`, `idpic1`, `idpic2`, `pimg1`, `pimg2`, `pimg3`, `status`, `im_money`, `tel_money`, `create_time`, `update_time`)
VALUES
	(1,'13245675486',NULL,'','',NULL,'','',0,'d296507972d67a50fd62060b637d1605','5mFIIg',NULL,NULL,NULL,NULL,NULL,0,0.00,0.00,'2019-12-05 14:02:16','2019-12-22 10:41:58'),
	(2,'13245675487',NULL,'','',NULL,'','',0,'5026692e16d7ee025da1c6157f5974e6','M7vokL',NULL,NULL,NULL,NULL,NULL,0,0.00,0.00,'2019-12-16 13:51:30',NULL);

/*!40000 ALTER TABLE `t_zjuser` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_zxuser
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_zxuser`;

CREATE TABLE `t_zxuser` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `source_id` varchar(50) DEFAULT '' COMMENT '注册ID',
  `source_type` tinyint(1) DEFAULT '0' COMMENT '注册类别',
  `recommender` varchar(20) DEFAULT '' COMMENT '推荐人',
  `nick_name` varchar(50) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(50) DEFAULT '' COMMENT '头像',
  `zj_id` bigint(20) DEFAULT NULL COMMENT '专家ID',
  `status` tinyint(11) DEFAULT '0' COMMENT '状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='咨询用户';




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
