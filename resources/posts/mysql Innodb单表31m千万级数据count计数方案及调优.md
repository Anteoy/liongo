---
date: 2017-11-28 15:26:00
title: mysql Innodb单表31m千万级数据count计数方案及调优
categories:
    - mysql
tags:
    - mysql
---

# ENV
1. 线上环境为RDS,版本5.7.15
    ```
    select version()
    output:
    5.7.15-log
    ```
2. 测试环境为docker搭建的mysql,版本5.7.19
    ```
    select version()
    output:
    5.7.19
    ```
3. 单表3000万+的class表以及20万+的学校表,需要使用count查询实时数量用于分页,延迟不能太高,否则影响业务
4. 因需要使用事务功能,使用存储引擎为Innodb(MyISAM count是自动计数单独保存,Innodb需要每次扫描表进行统计)
5. 本文使用class表进行示例表述,school同理

# OPTIMIZE
## 出现的第一个问题是RDS线上mysql的查询速度始终没有测试库的快,相同的数据和存储结构,索引数据都相同(一开始线上使用count完全不能查询,会出现等待超时).

1. 查看索引
```
show index from consumer.class;
output:
'class', '0', 'PRIMARY', '1', 'id', 'A', '28663646', NULL, NULL, '', 'BTREE', '', ''
'class', '0', 'UQE_class_loginName', '1', 'loginName', 'A', '28663646', NULL, NULL, 'YES', 'BTREE', '', ''
'class', '1', 'IDX_class_school_id', '1', 'school_id', 'A', '211268', NULL, NULL, '', 'BTREE', '', ''
'class', '1', 'grade_id', '1', 'grade_id', 'A', '8644', NULL, NULL, 'YES', 'BTREE', '', ''
'class', '1', 'schuid', '1', 'schuid', 'A', '216557', NULL, NULL, 'YES', 'BTREE', '', ''
```
2. 测试时间
    ```
        set profiling = 1;
        SELECT count(*) FROM consumer.class;
        show profiles;
    ```
3. 分别分析sql的执行
    ```
        explain select count(*) from consumer.class ;
        test output:
        '1', 'SIMPLE', 'class', NULL, 'index', NULL, 'IDX_class_school_id', '4', NULL, '28663646', '100.00', 'Using index'
        online output:
        1
        1
        SIMPLE
        null
        null
        null
        null
        null
        null
        null
        null
        null
        Select tables optimized away
    ```
    发现线上版本的mysql是经过自己编译器优化的Select tables optimized away,但是效率确实低到不能接受(单独这样查询几分钟过后仍然查不出来,并且显示超时),这是因为mysql5.7.*版本机制相关的问题,具体可参考:
    [https://bugs.mysql.com/bug.php?id=80580](https://bugs.mysql.com/bug.php?id=80580)
    [https://stackoverflow.com/questions/27377549/select-count-not-using-index](https://stackoverflow.com/questions/27377549/select-count-not-using-index)
5. 强制使用索引,解决了上面线上查询几分钟仍不能查询到结果后返回超时的问题
    ```
        select count(`id`) from consumer.class force index(primary) where id > 0
        explain select count(`id`) from consumer.class force index(primary) where id > 0
    ```
    但是线上RDS速度仍然不尽人意,需要大概2分多钟,而测试自建的在10-20秒,目前猜测是因为线上RDS使用的是机械硬盘,而我们测试环境自建的为SSD.因为是阿里云的RDS,所以不知道内部做了些什么处理.
6. 使用mysql 事件和触发器解决,定时执行事件event进行统计并插入一个专门用于记录的统计表,然后触发器监听指定表的insert和delete操作,分别对记录进行加1和减1.
可参考我如下sql创建:
    ```
        use consumer;
        CREATE TABLE `class_count` (
          `key` varchar(50) NOT NULL,
          `value` varchar(100) NOT NULL,
          PRIMARY KEY (`key`)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
        
        CREATE TABLE `school_count` (
          `key` varchar(50) NOT NULL,
          `value` varchar(100) NOT NULL,
          PRIMARY KEY (`key`)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
        
        drop event `consumer`.`update_class_count_1`;
        CREATE EVENT `consumer`.`update_class_count_1` 
          ON SCHEDULE EVERY 10 day
          STARTS  '2017-09-22 01:00:00'  ON COMPLETION NOT PRESERVE  
          ENABLE  
          DO INSERT INTO consumer.class_count (`key`, `value`)
          VALUES ('cal_count',(select count(*) from consumer.class force INDEX (IDX_class_school_id) where `school_id` != 0))
          ON DUPLICATE KEY UPDATE value=VALUES(value);
          
          
          drop event `consumer`.`update_school_count`;
        CREATE EVENT `consumer`.`update_school_count` 
          ON SCHEDULE EVERY 10 DAY
          STARTS  '2017-09-22 18:16:04'  ON COMPLETION NOT PRESERVE  
          ENABLE  
          DO INSERT INTO consumer.school_count (`key`, `value`)
          VALUES ('cal_count', (select count(`id`) from consumer.school_login force index(primary) where id > 0))
          ON DUPLICATE KEY UPDATE value=VALUES(value);
          
          
          
          DROP TRIGGER IF EXISTS `consumer`.`count_down_class`; 
        CREATE  TRIGGER `consumer`.`count_down_class` 
        AFTER DELETE  ON consumer.class
         FOR EACH ROW  
        UPDATE `class_count`
        SET 
          `class_count`.`value` = `class_count`.`value` - 1 
        WHERE
          `class_count`.`key` = "cal_count"; 
          
          
          DROP TRIGGER IF EXISTS `consumer`.`count_down_school`; 
        CREATE  TRIGGER `consumer`.`count_down_school` 
        AFTER DELETE  ON consumer.school_login
         FOR EACH ROW  
        UPDATE `school_count`
        SET 
          `school_count`.`value` = `school_count`.`value` - 1 
        WHERE
          `school_count`.`key` = "cal_count"; 
          
          
          
          DROP TRIGGER IF EXISTS `consumer`.`count_up_class`; 
        CREATE  TRIGGER `consumer`.`count_up_class` 
        AFTER INSERT  ON consumer.class
         FOR EACH ROW  
        UPDATE `class_count`
        SET 
          `class_count`.`value` = `class_count`.`value` + 1 
        WHERE
          `class_count`.`key` = "cal_count"; 
        
        
        DROP TRIGGER IF EXISTS `consumer`.`count_up_school`; 
        CREATE  TRIGGER `consumer`.`count_up_school` 
        AFTER INSERT  ON consumer.school_login
         FOR EACH ROW  
        UPDATE `school_count`
        SET 
          `school_count`.`value` = `school_count`.`value` + 1 
        WHERE
          `school_count`.`key` = "cal_count"; 
    ```
    这里线上RDS特别需要注意,我也是看了好久,因为mysql开启event必须设置SET GLOBAL event_scheduler = 1; 开启事件功能,否则就算创建了也不会生效,而RDS执行这个sql指令的时候,一直提示我没有super权限(测试库自建的mysql可以直接设置),但是我确实是使用root用户登录的,后来在DMS的RDS的管理界面看到有此参数配置,手动修改下就生效了.完成后可以使用SHOW VARIABLES LIKE 'event_scheduler',需要看到'event_scheduler', 'ON',最后检查下events和triggers有没有创建好
    ```
    SHOW TRIGGERS;
    SHOW EVENTS;
    ```
    另外,count(*),count(1) 和 count([column])我这里经过测试,两者执行时间基本一致,查阅相关资料发现mysql解释器会自动优化count(*),所以从执行效率上说count(*),count(1)和count([column])实际执行效率几乎相同,都是经过了解释器优化的.而"InnoDB handles SELECT COUNT(*) and SELECT COUNT(1) operations in the same way. There is no performance difference.",即count(*) = count(1) 计数包括null,count([column])计数不包括null,可参阅mysql 5.6中文档:
    [https://dev.mysql.com/doc/refman/5.6/en/innodb-restrictions.html](https://dev.mysql.com/doc/refman/5.6/en/innodb-restrictions.html)
    [https://dev.mysql.com/doc/refman/5.6/en/group-by-functions.html#function_count](https://dev.mysql.com/doc/refman/5.6/en/group-by-functions.html#function_count)

    *注意:使用offset的时候,因为mysql不知道数据表的行数据是否连续,所以需要遍历数据表进行查找,offset的值越大,查询的速度耗时越长,我这里limit 20 offset 30579060,耗时约60-70秒,如果不要求精确度,可以用where id来进行获取,这样就不会吃哦那头遍历到offset 30579060,当然如果业务允许的情况,比如如果不使用innode的事务,更倾向使用MYISAM的话,将会是一个更好的解决方案*
    
# REFERENCES
1. [https://stackoverflow.com/questions/19267507/how-to-optimize-count-performance-on-innodb-by-using-index?noredirect=1&lq=1](https://stackoverflow.com/questions/19267507/how-to-optimize-count-performance-on-innodb-by-using-index?noredirect=1&lq=1)db-by-using-index?noredirect=1&lq=1)ing-index?noredirect=1&lq=1)