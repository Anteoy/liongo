---
date: 2017-01-27 16:03:00
title: mongo 3.0 备份和还原数据库
categories:
    - DB
tags:
---

### 主要记录下在mongo 3.0的操作
####备份示例
```
./mongodump -h localhost -d liongo -o ./

```
#####还原示例
错误方式：
```
./mongorestore -h 127.0.0.1 -d liongo --directoryperdb /home/zhoudazhuang/company-zhoudazhuang/liongo/note.bson
```
会报错：

```
2017-01-27T15:31:54.217+0800	error parsing command line options: --dbpath and related flags are not supported in 3.0 tools.
See http://dochub.mongodb.org/core/tools-dbpath-deprecated for more information
2017-01-27T15:31:54.217+0800	try 'mongorestore --help' for more information
```
使用mongorestore --help,正确还原方式为（去掉--directoryperdb）：

```
./mongorestore -h 127.0.0.1 -d liongo /home/zhoudazhuang/company-zhoudazhuang/liongo/note.bson --drop
```
