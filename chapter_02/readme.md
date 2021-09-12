### 作业说明

* 作业在main.go 中，用TODO注释写了解释说明
* 代码可运行，数据库自己搭在阿里云的docker镜像

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