---
date: 2017-08-31 21:22:00
title: Rancher部署kubernetes爬坑记
categories:
    - kubernetes
tags:
    - Rancher,kubernetes
---

### 引言：
　本文不会记录详细的部署过程，仅记录下使用Rancher部署kubernetes踩过的几个小坑，如果你需要详细的部署过程,可以参考[此处官方教程](http://www.cnrancher.com/rancher-k8s-accelerate-installation-document/)，这里面有详细的部署说明，另外可同时参考下此文，也许对你在部署中遇到的问题有所帮助。由于个人能力有限，如有不当之处，欢迎指正。
　
### 环境
   1. 一台ubuntu 16.04 服务器 作为Rancher Server 宿主机，并且加入k8s集群
   2. 一台virtualbox 虚拟的centos 7.3 服务器，并且加入k8s集群
　
### localhost 问题
问题描述：在安装好Rancher server，以及新建k8s环境后，需要在k8s中添加主机节点host，按照描述在需要加入的节点中运行，结果报错:
```
requests.exceptions.ConnectionError: HTTPConnectionPool(host='localhost', port=8080): Max retries exceeded with url: /v1 (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f945c7bee50>: Failed to establish a new connection: [Errno 111] Connection refused',))
```
原因：因为默认是在localhost里部署的rancher server,导致rancher默认配置的地址为localhost，而不是局域网或公网ip地址
解决：需要进入rancher的dashboard ui界面，进入设置界面，从已经选择的localhost切换为ip地址的服务.如图：
![settings](http://img.blog.csdn.net/20170831215857509?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvWWFuX0Nob3U=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
延伸：感谢nevotheless [https://github.com/rancher/rancher/issues/6623](https://github.com/rancher/rancher/issues/6623)

### docker 版本问题
问题描述：在k8s中添加主机节点host完成后，在节点上提示不支持的docker版本，进入查看，发现同时支持rancher和k8s的docker版本为1.12.3以及1.12这个大版本号以后的版本，我节点的docker是17.03.1-ce版本的，友情提示，我曾尝试过1.12.6，原本看文档以为是支持的，结果1.12.6并不支持，1.12.6就不要使用了。在降级的时候，由于网络等原因，折腾了不少时间，主要在purge,autoremove,remove docker engine后，重新安装1.12.3版本，总是启动失败，安装完成后报错信息为：
```
Job for docker.service failed because the control process exited with error code. See "systemctl status docker.service" and "journalctl -xe" for details.
```
原因：systemctl status docker.service 查看部分日志如下：
```
level=fatal msg="Error creating cluster component: error while loading TLS Certificate in /var/lib/docker/swarm/certificates/swarm-node.crt: x509: certificate has expired or is not yet valid"
```
开始没有仔细看，这也算给自己长了个教训，最后发现有一句"Error creating cluster component: error while loading TLS Certificate in /var/lib/docker/swarm/certificates/swarm-node.crt: x509: certificate has expired or is not yet valid",发现是以前安装的swarm的证书过期或无效了
解决：将/var/lib/docker/swarm目录删除或更名.重启服务.

延伸：感谢Jason-ZH的分享 [https://my.oschina.net/JasonZhang/blog/820786](https://my.oschina.net/JasonZhang/blog/820786)