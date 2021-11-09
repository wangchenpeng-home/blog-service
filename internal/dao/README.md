数据访问层（Database Access Object）

CREATE DATABASE
IF
    NOT EXISTS blog_service DEFAULT CHARACTER
    SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;


#更新环境变量

export PATH=$PATH:/usr/local/mysql/bin
source /etc/profile

#创建标签页

CREATE TABLE `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '标签名称',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

CREATE TABLE `blog_article` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`title` varchar(100) DEFAULT '' COMMENT '文章标题',
`desc` varchar(255) DEFAULT '' COMMENT '文章简述',
`cover_image_url` varchar(1255) DEFAULT '' COMMENT '封面地址图片',
`content` longtext COMMENT '文章内容',
`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1为启用',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';


CREATE TABLE `blog_article_tag` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`article_id` int(11) NOT NULL COMMENT '文章ID',
`tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';

#创建认证表
CREATE TABLE `blog_auth` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`app_key` varchar(20) DEFAULT '' COMMENT 'KEY',
`app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
`created_by` varchar(100) DEFAULT '' COMMENT '创建者',
`modified_by` varchar(100) DEFAULT '' COMMENT '修改者',
`created_on` int(32),
`modified_on` int(32),
`deleted_on` int(32),
`is_del` int(8),
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

#插入数据
INSERT INTO `blog_service`.`blog_auth`(
`id`,`app_key`,`app_secret`,`created_on`,`created_by`,
`modified_on`,`deleted_on`,`is_del`)
) VALUES (1, 'eddycjy', 'wangchenpeng-home', 0, 'eddycjy', 0, 0, 0);
