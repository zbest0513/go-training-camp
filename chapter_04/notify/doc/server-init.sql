CREATE TABLE `notify_user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uuid` varchar(40) NOT NULL COMMENT 'uuid',
    `name` varchar(100) DEFAULT '' COMMENT '用户名',
    `mobile` char(11) DEFAULT '' COMMENT '手机号',
    `email` varchar(100) DEFAULT '' COMMENT '邮箱',
    `status` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='用户';

CREATE TABLE `notify_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uuid` varchar(40) NOT NULL COMMENT 'uuid',
    `name` varchar(100) DEFAULT '' COMMENT '用户名',
    `desc` varchar(300) DEFAULT '' COMMENT '描述',
    `status` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签';

CREATE TABLE `notify_template` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uuid` varchar(40) NOT NULL COMMENT 'uuid',
    `name` varchar(100) DEFAULT '' COMMENT '模版名称',
    `desc` varchar(300) DEFAULT '' COMMENT '描述',
    `content` text COMMENT '内容',
    `status` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='模版';

CREATE TABLE `notify_tag_template` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `template_uuid` varchar(40) NOT NULL COMMENT '模版uuid',
    `tag_uuid` varchar(40) NOT NULL COMMENT '标签uuid',
    `status` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_template_tag` (`template_uuid`,`tag_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签模版关系表';

CREATE TABLE `notify_user_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_uuid` varchar(40) NOT NULL COMMENT '用户uuid',
    `tag_uuid` varchar(40) NOT NULL COMMENT '标签uuid',
    `status` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_user_tag` (`user_uuid`,`tag_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户标签关系表';