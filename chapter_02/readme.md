### 初始化sql 

```
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '用户名',
  `age` int(20) DEFAULT '0' COMMENT '年龄',
  `card` varchar(100) DEFAULT '' COMMENT '银行卡',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='用户';
```