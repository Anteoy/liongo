---
date: 2016-12-28 23:08:00
title: ubuntu16.04 自动挂载新分区到指定目录下
categories:
    - linux
tags:
    - linux
---

###引言：
　　起因：ubuntu系统使用固态硬盘120G不够使用，如今已无法满足日常需要，于是增加了一枚机械硬盘，分别分了两个空闲分区，依次使系统自启时挂在到/home/和/usr/指定目录下
###安装环境：
 1. ubuntu 16.04 LTS
 2. 有剩余未分配空间的硬盘
###安装过程：
1. 查看硬盘所有分区并记录下待挂在分区（主要以硬盘，分区大小和格式确定）
	```
	fdisk -l
	```
2. 格式化分区为ext4
	```
	mkfs.ext4 /dev/sda7
	```
3. 查看分区UUID
	```
	sudo blkid
	```
4. 编辑系统挂载配置文件/etc/fstab
	```
	vim /etc/fstab
	```
	增加：
	```
	# new /dev/sda6
	UUID=fdccb3ee-1efb-4822-b947-9ad28a4c4243 /home/zhoudazhuang/dev_sda_6 ext4 defaults 0	0
	# new /dev/sda7
	UUID=f3ff6a9b-71d6-4daa-bfc2-805d50a03c50 /usr/dev_sda_7 ext4 defaults 0 0
	```
	注：格式为 设备名称 挂载点 分区类型 挂载选项 dump选项 fsck选项
	dump选项--这一项为0，就表示从不备份。如果上次用dump备份，将显示备份至今的天数。
	fsck选项 --启动时fsck检查的顺序。为0就表示不检查，（/）分区永远都是1，其它的分区只能从2开始，当数字相同就同时检查（但不能有两1）
4. 修改文件所属权限
	```
	sudo chown -R zhoudazhuang:zhoudazhuang /home/zhoudazhuang/dev_sda_6
	sudo chmod -R 4755 /home/zhoudazhuang/dev_sda_6
	```
	在注：chmod使用的数字的意思： 读（r=4），写（w=2），执行（x=1）可读可写为4+2=6 依次内存 755表示的是文件所有者权限7（三者权限之和），与所有者同组用户5（读+执行），其他用户同前一个5，这里的4的意思是（其他）用户执行拥有所有者相同的文件权限（对于其他要使用的文件）
5. 重启
```
sudo reboot now
```