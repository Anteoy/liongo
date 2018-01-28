---
date: 2018-01-14 23:33:00
title: 使用 let's encrypt certbot部署https网站
categories:
    - 架构
tags:
    - let's encrypt,certbot
---

### 前言
let's encrypt 是免费的ssl/tls 证书颁发的机构，致力于实现整个Web的TLS/SSL认证。https可降低网站被劫持的风险，并具有更好的加密性能，避免用户信息泄露，增强网站的安全性。
### 准备
1. 已解析正确的域名 www.anteoy.me
2. A记录所指向的服务器
3. nginx
### 环境
1. GCE ubuntu16.04
### let's encrypt认证过程
certbot是let's encrypt的官方客户端工具，客户端在认证过程中会在host上生成自己的加密文件，let's encrypt服务端访问客户端提供的域名并尝试去获取这个文件，如果成功获取并确认是客户端生成的正确文件，则确认客户端所在主机的域名控制权，然后开始为此域名颁发CA证书。

#### 部署过程
1. 使用nginx以便let's encrypt验证域名 
    * 安装nginx
    ```
    sudo apt-get install nginx
    ```
    * 编辑配置文件，配置nginx的webroot目录和需要授权访问的隐藏目录.well-known
    
        ```
        vi /etc/nginx/conf.d/cert.conf
        ```
        ```
        server {
            listen       80;
            root /usr/share/nginx/html;
            server_name www.anteoy.me;
            location ~ /.well-known {
              allow all;
            }
        }
        ```
    * 平滑重启nginx服务器
        ```
            nginx -s reload
        ```
2. 生成CA
    ```
        sudo certbot certonly --webroot -w /usr/share/nginx/html/ -d www.anteoy.me --agree-tos --email anteoy@gmail.com
    ```
    使用 --webroot 模式会在 /var/www/example 中创建 .well-known 文件夹，这个文件夹里面包含了一些验证文件，letsencrypt 会通过访问 www.anteoy.me/.well-known/acme-challenge 来验证此域名是否绑定的这个服务器。--agree-tos 参数是你同意他们的协议。注意certbot提供了两种CA生成方式，其中一种是certbot提供服务器而不是像我们这里使用nginx，此certbot的--standalone 模式会自动启用服务器的 443 端口，来验证域名的归属。如果我们本生有服务占用了443和80,则必须先关掉。推荐使用--webroot方式,执行正确部分响应如下：
    ```
        IMPORTANT NOTES:
         - Congratulations! Your certificate and chain have been saved at:
           /etc/letsencrypt/live/www.anteoy.me/fullchain.pem
           Your key file has been saved at:
           /etc/letsencrypt/live/www.anteoy.me/privkey.pem
           Your cert will expire on 2018-04-14. To obtain a new or tweaked
           version of this certificate in the future, simply run certbot
           again. To non-interactively renew *all* of your certificates, run
           "certbot renew"
         - If you like Certbot, please consider supporting our work by:
           Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
           Donating to EFF:                    https://eff.org/donate-le
    ```
3. 利用nginx部署https
在/etc/nginx/conf.d/下增加两个server配置文件，以*.conf命名分别配置两个server
    ```
        server {
            listen 443 ssl;
            server_name www.anteoy.me;
            ssl_certificate /etc/letsencrypt/live/www.anteoy.me/fullchain.pem;
            ssl_certificate_key  /etc/letsencrypt/live/www.anteoy.me/privkey.pem;
            ssl_trusted_certificate /etc/letsencrypt/live/www.anteoy.me/chain.pem;
            root /usr/share/nginx/html;
            location / {
              proxy_pass http://web服务的ip或者域名:8080/;
            }
        }
    ```
    ```
        server {
            listen       80;
            server_name www.anteoy.me;
            return 301 https://$host$request_uri;
        }
    ```
    listen 80端口主要是为了在用户访问网站的时候未输入https，使用http的方式访问80,则自动跳转请求https的访问地址
    重启nginx：
    ```
        nginx -s reload
    ```
4. 配置crontab
由于let's encrypt 生成的CA有效时间只有3个月，所以在CA到期以后我们需要手动进行更新，重新获取，或者使用Linux的crontab定时任务定时获取
首先完成步骤3后检测能否正常更新证书：
    ```
        certbot renew --dry-run
    ```
    然后编辑自定义脚本regen.sh：
    ```
        #!/bin/bash
        # 续签
        /usr/bin/certbot renew --quiet
        # 重启 nginx
        /usr/sbin/nginx -s reload
    
    ```
    查看任务列表
    ```
        crontab -l
    ```
    增加cron
    ```
        crontab -e
    ```
    注意如果是首次添加则会选择编辑器，按找自己习惯选择就行，我这里选择的是vi
    在文件末尾追加：
    ```
        00 03 1 * * /youpath/regen.sh
    ```
    执行此脚本测试是否正常：
    ```
        chmod +x regen.sh
        ./regin.sh
    ```
    重启crontab
    ```
        sudo systemctl restart cron
    ```
    
### 参考
1. [https://certbot.eff.org/](https://certbot.eff.org/)
2. [https://blog.guorenxi.com/43.html](https://blog.guorenxi.com/43.html)
3. [https://segmentfault.com/a/1190000005797776](https://segmentfault.com/a/1190000005797776)
4. [https://linuxstory.org/deploy-lets-encrypt-ssl-certificate-with-certbot/](https://linuxstory.org/deploy-lets-encrypt-ssl-certificate-with-certbot/)-encrypt-ssl-certificate-with-certbot/)t/)