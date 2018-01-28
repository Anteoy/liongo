---
date: 2017-10-20 11:00:00
title: squid stunnel 为 docker 配置代理服务器
categories:
    - linux,docker
tags:
    - linux,squid,stunnel
---

### 目地
为k8s的docker服务提供http/https代理,解决docker无法pull gcr.io/google_containers 谷歌镜像问题

### 环境 
1. GCE ubuntu 16.04 
2. k8s集群机器 ubuntu16.04

#### 简要步骤

* GCE 搭建squid正向http/https代理服务器
1. 直接使用apt-get install 安装
    ```
    apt-get install squid3 -y
    ```
    *注意:配置文件在/etc/squid或/etc/squid3下,根据系统不同可能会有一点差异,由于这里进行快速安装,不需要暴露端口给外部使用,也不需要密码,所以配置文件我这里保持默认*
    
* GCE 安装stunnel代理服务器
1. stunnel主要用来在GCE和k8s机器上代理的数据传输进行加密,否则明文传输很快会被GFW拦截.注意stunnel分为服务端和客户端,GCE上安装服务端,在k8s集群上安装客户端
2. 直接使用apt-get install 安装服务端
    ```
    apt-get install stunnel4 -y
    ```
3. 编辑配置文件  vim /etc/stunnel/stunnel.conf
    ```
    client = no #是否为客户端 这里是服务端填写no
    [squid]
    accept = 65501
    connect = 127.0.0.1:3128 #本地squid服务地址
    cert = /etc/stunnel/stunnel.pem #下一步生成的证书地址
    ```
4. openssl生成证书,用户stunnel加密解密
    ```
    openssl genrsa -out key.pem 2048
    openssl req -new -x509 -key key.pem -out cert.pem -days 1095
    cat key.pem cert.pem >> /etc/stunnel/stunnel.pem
    ```
    *注意：创建证书时，系统会要求您提供一些国家/地区信息，可随便输入，但是当被要求输入“Common Name”时，您必须输入正确的hostname或IP地址（VPS）,我这里输入的ip地址。*
5. 通过配置/etc/default/stunnel4文件启用自动启动，vim /etc/default/stunnel4
    ```
    #将ENABLED更改为1：
    
    ENABLED=1
    ```
6. 重新启动Stunnel使配置生效，使用以下命令：
    ```
    /etc/init.d/stunnel4 restart
    ```
    
* K8S 集群机器分别搭建stunnel
1. 安装步骤几乎和上面相同
2. scp或其他方法把证书拷贝到k8s集群中
2. 配置文件不同
    ```
    cert = /etc/stunnel/stunnel.pem #和服务端完全相同的证书
    client = yes #声明为客户端
    [squid]
    accept = 127.0.0.1:65502 #本地代理的端口,即为http/https代理地址
    connect = {GCE_IP}:65501 #GCE 服务端ip和端口
    ```
    
四. 浏览器SwitchyOmega代理穿透GFW
1. 配置添加http代理127.0.0.1:65502 即可
四. docker添加http/https代理请参考官方文档
[https://docs.docker.com/engine/admin/systemd/#start-manually](https://docs.docker.com/engine/admin/systemd/#start-manually)
### 参考文献
[https://www.digitalocean.com/community/tutorials/how-to-set-up-an-ssl-tunnel-using-stunnel-on-ubuntu](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-ssl-tunnel-using-stunnel-on-ubuntu)-ubuntu)