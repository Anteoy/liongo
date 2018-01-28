---
date: 2017-09-21 22:33:00
title: archlinux下网易云音乐netease-cloud-music部分问题
categories:
    - linux
tags:
    - linux,archlinux,网易云音乐
---

*个人为网易云音乐重度用户，从ubuntu到arch后，感觉netease-cloud-music仍然在ubuntu中支持更好，毕竟linux下的netease-cloud-music是由网易云音乐和深度deepin联合开发的，ubuntu有官方的deb包，arch的aur中的netease-cloud-music是基于此deb进行打包的，在体验上没有差别，在arch中用起来感觉也还不错，在此记录下使用过程中遇到的问题以及解决方式，本文会不定时更新*
1. 启动报错，客户端无法打开

这个问题具体报错记不清了，问题为点击无法打开，命令行运行报错，后来在AUR上看到，在命令行上使用 --no-sandbox 关闭沙箱运行，则可正常使用
2. 这个问题有点坑，在用了一段时间后，发现网易云音乐再次无法打开，报错如下：
```
[0921/220732:ERROR:browser_main_loop.cc(203)] Running without the SUID sandbox! See https://code.google.com/p/chromium/wiki/LinuxSUIDSandboxDevelopment for more information on developing with the sandbox on.
```
后来验证这个Log并不是引起客户端无法启动的原因，找了蛮久原因，期间曾尝试使用wine，web版本等，但确实还是用起来不习惯，最后使用ps -ef | grep netease 发现系统启动了好几个相关进程（linux没有线程，都是用进程模仿的），如下：
```
➜  ~ sudo ps -ef | grep netease
[sudo] zhoudazhuang 的密码：
zhoudaz+ 11960 11896  0 21:37 pts/0    00:00:00 /bin/sh /usr/bin/netease-cloud-music --no-sandbox
zhoudaz+ 11961 11960  3 21:37 pts/0    00:01:16 /usr/lib/netease-cloud-music/netease-cloud-music --no-sandbox
zhoudaz+ 11965 11961  0 21:37 pts/0    00:00:00 /usr/lib/netease-cloud-music/netease-cloud-music --type=zygote --no-sandbox --lang=en-US --log-file=/home/zhoudazhuang/.cache/netease-cloud-music/Cef/console.log --log-severity=disable
zhoudaz+ 11991 11965  1 21:37 pts/0    00:00:23 /usr/lib/netease-cloud-music/netease-cloud-music --type=zygote --no-sandbox --lang=en-US --log-file=/home/zhoudazhuang/.cache/netease-cloud-music/Cef/console.log --log-severity=disable
zhoudaz+ 13325 13293  0 22:12 pts/7    00:00:00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn netease
```
然后尝试杀死kill -9 [pid] ，再用netease-cloud-music --no-sandbox启动，则可以正常使用了（依然存在ERROR:browser_main_loop，可忽略）。

此问题原因：因为个人桌面环境为xfce，在ubuntu下关机习惯直接使用shutdown now直接关闭系统（未关闭打开的程序，比如网易云音乐客户端），然后系统在下一次启动的使用会自动帮你启动部分程序（比如网易云音乐客户端，chrome不会），而系统帮你启动的时候使用的是netease-cloud-music命令（没有--no-sandbox关闭沙箱环境），导致进程已启动，所以后来如何启动都无法再次打开启动了，并且也没有额外报错信息。

验证如下，直接关机然后开机，查看相关进程是否已启动：
```
sudo ps -ef | grep netease
```
结果：
```
➜  ~ sudo ps -ef | grep netease
[sudo] zhoudazhuang 的密码：
zhoudaz+  1133  1076  1 22:19 tty2     00:00:00 /usr/lib/netease-cloud-music/netease-cloud-music -session 267fbbadc-1458-4223-8e6a-3a9d82521207_1506003342_809192
zhoudaz+  1307  1133  0 22:19 tty2     00:00:00 /usr/lib/netease-cloud-music/chrome-sandbox /usr/lib/netease-cloud-music/netease-cloud-music --type=zygote --lang=en-US --log-file=/home/zhoudazhuang/.cache/netease-cloud-music/Cef/console.log --log-severity=disable
zhoudaz+  1312  1307  0 22:19 tty2     00:00:00 /usr/lib/netease-cloud-music/netease-cloud-music --type=zygote --lang=en-US --log-file=/home/zhoudazhuang/.cache/netease-cloud-music/Cef/console.log --log-severity=disable
zhoudaz+  1341  1312  0 22:19 tty2     00:00:00 /usr/lib/netease-cloud-music/netease-cloud-music --type=zygote --lang=en-US --log-file=/home/zhoudazhuang/.cache/netease-cloud-music/Cef/console.log --log-severity=disable
zhoudaz+  1521  1171  0 22:19 pts/0    00:00:00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn netease
```
确实是由此原因引起的。