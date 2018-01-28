---
date: 2016-12-03 18:29:00
title: jottings-ubuntu16.04 lts的完整克隆
categories:
    - linux
tags:
---

**安装环境：**


 1. ubuntu 16.04 LTS

 2. 一台待克隆的计算机，一台原始ubuntu16.04计算机

**安装过程：**
1. 在新的计算机中安装好ubuntu16.04 LTS（这篇文章主要介绍系统的克隆，对于初始安装这里不再赘述,所有操作请确保自己拥有root权限）
2. 使用tar压缩源计算机的相关文件（包括配置，软件，文件等）
`tar -zcpPf /media/zhoudazhuang/NEWSMY/ubuntu.tar.gz --exclude=/proc --exclude=/lost+found --exclude=/media/zhoudazhuang/NEWSMY/ubuntu.tar.gz --exclude=/media --exclude=/mnt --exclude=/sys  / >/dev/nulll`
 *注：f参数需要放在最后面，否则会无法操作，在进行解压操作的时候也是如此，使用>/dev/null表示把标准输出放入黑洞，即不会输出标准输出，这样能清晰地看到控制台的标准错误输出*
	- z 使用gzip压缩
	- c create创建一个文件
	- p 保持原文件的属性
	- P 使用绝对路径压缩
	- f 使用档名
	- --exclude 为排除文件，对此部分的文件不进行压缩
3. 备份新装系统的/boot和/etc/fstab文件夹，/boot用于记录各自电脑的启动，如grub信息，/etc/fstab记录分区的挂在信息，这两个必须保留，否则开机无法进入系统。
    `tar -zcvpf bootandfstab.tar.gz /boot /etc/fstab`
4. tar 解压ubuntu.tar.gz到新系统/目录

    `tar -zxvpf ubuntu.tar.gz -C /`

    -C表示解压使用绝对路径
5. 使用上一步的方法还原之前备份的/boot和/etc/fstab目录

    `tar -zxvpf bootandfstab.tar.gz -C /`
**后记**
1. 在使用tar压缩和解压系统压缩包的时候，总是报错
	`tar: 由于前次错误，将以上次的错误状态退出`
	最开始使用-v参数在控制台打印大量日志信息，于是后来使用>/dev/null屏蔽了标准输出，只打印标准错误输出，原本以为是因为部分文件权限问题（虽然我使用的是root用户），但屏蔽标准输出后仍然无法定位错误原因，错误日志依旧，而压缩文件亦是可以使用的，系统也和源系统一致，于是暂未深究，等有时间再仔细看看。
2. tar指令参数中，我曾把-p放在-f之后，但这样使用会直接报错，无法继续，查阅-f参数为-f<备份文件>或--file=<备份文件>：指定备份文件；-p为保持原文件属性（不随用户变动而改变），这里暂且理解为一种语法约定。
