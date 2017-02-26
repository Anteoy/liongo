---
date: 2016-09-18 20:35:00
title: java.util.Date,java.sql.Date及java.sql.Timestamp
categories:
    - java
tags:
---

java.sql.Date及java.sql.Timestamp继承自java.util.Date ，三个类都可以使用getTime()进行互换，java.util.Date有无参构造方法获取当前时间，其余两个没有。Timestamp为时间戳，和sql.Date的精确度一样，但表示当前时间更加方便（另外在hibernate中使用idea自动生成表的pojo时，会把sql.Date写为Timestamp）,部分示例如下：

```
package test;

import java.sql.Date;
import java.sql.Timestamp;
import java.text.ParseException;
import java.text.SimpleDateFormat;

/**
 * Created by zhoudazhuang on 2016/7/23.
 */
public class Test{

    public static void main(String[] args) throws ParseException {
        System.out.println("hello world");
        Timestamp timestamp = new Timestamp(new java.util.Date().getTime());
        Date date = new Date(new java.util.Date().getTime());
        java.util.Date date1 = new java.util.Date();
        java.util.Date date2 = new java.util.Date(timestamp.getTime());
        java.util.Date date3 = new java.util.Date(date.getTime());
        System.out.println("timestamp:"+timestamp);
        System.out.println("sql:"+date);
        System.out.println("util date1:"+date1);
        System.out.println("util date2:"+date2);
        System.out.println("util date3:"+date3);
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd hh:mm:ss");

        String str = simpleDateFormat.format(date2);
        System.out.println("yyyy-MM-dd hh:mm:ss"+str);
        java.util.Date date4 = simpleDateFormat.parse(str);
        System.out.println("util date4:"+date4);
        SimpleDateFormat simpleDateFormat1 = new SimpleDateFormat("yyyy");
        System.out.println("yyyy:"+simpleDateFormat1.format(date2));

        SimpleDateFormat simpleDateFormat2 = new SimpleDateFormat("MM");
        System.out.println("MM:"+simpleDateFormat2.format(date2));

    }

}
```

结果为：

```
hello world
timestamp:2016-09-18 20:33:29.328
sql:2016-09-18
util date1:Sun Sep 18 20:33:29 CST 2016
util date2:Sun Sep 18 20:33:29 CST 2016
util date3:Sun Sep 18 20:33:29 CST 2016
yyyy-MM-dd hh:mm:ss2016-09-18 08:33:29
util date4:Sun Sep 18 08:33:29 CST 2016
yyyy:2016
MM:09
```
