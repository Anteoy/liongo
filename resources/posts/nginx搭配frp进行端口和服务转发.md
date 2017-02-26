---
date: 2017-02-20 16:06:00
title: nginx搭配frp进行端口和服务转发
categories:
    - 架构
tags:
---

###引言：
　　java接入三方运营商服务接口，需要可供回调的公网接口，并在本地两台（及以上）调试接口服务，于是使用开源frp进行穿透，nginx搭配负责分发请求到不同机器。

###安装环境：
 1. ubuntu 16.04 LTS 一台服务器 两台客户机
 2. 已安装好nginx
###frp配置过程：

    参考我之前博文[http://blog.csdn.net/yan_chou/article/details/53406095](http://blog.csdn.net/yan_chou/article/details/53406095) 并分端口配置两份

 ###一台客户机中nginx配置过程：

修改nginx.conf

```


sudo vim /etc/nginx/nginx.conf

```

修改http中server节点为
```

server {

                server_name 127.0.0.1;#非必须

                listen 10000;#监听端口



                location / {

                proxy_pass http://172.16.64.16:8080/crawlerCallBack;//转发到另外一台客户机接口

                }

        }

```
###后记
*注：frp中指定不同端口请修改frps.ini中vhost_http_port,nginx中指定端口为listen*




