---
date: 2018-01-09 16:24:00
title: 使用nginx解决k8s traefik中basic auth的跨域问题
categories:
    - 运维
tags:
    - 运维,nginx,k8s,traefik,basic auth
---

### 目地
目前k8s ingress是配合traefik使用的，此时需要对某一个域名添加一个basic auth安全认证，原本traefik也可以正常配置（生产环境已有不少使用traefik basic auth），但是由于此处的域名需要在其他web域中调用，涉及到跨域问题，参考traefik文档未发现在k8s有关联说明解决basic auth相关跨域问题。后来分析了下nginx下的basic auth,最终使用nginx + ingress + traefik解决了这一问题。

#### 简要步骤
1. 生成basic auth用户密码文件
```
htpasswd -bc ngauth username password
```
2. 配置nginx的k8s configMap:
```
kind: ConfigMap
apiVersion: v1
metadata:
  name: nginx-conf
  namespace: kube-apps
data:
  nginx.conf: |
    user  nginx;
    worker_processes  1;

    error_log  /var/log/nginx/error.log warn;
    pid        /var/run/nginx.pid;


    events {
        worker_connections  1024;
    }


    http {
        default_type  application/octet-stream;

        log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                          '$status $body_bytes_sent "$http_referer" '
                          '"$http_user_agent" "$http_x_forwarded_for"';

        access_log  /var/log/nginx/access.log  main;

        sendfile        on;
        #tcp_nopush     on;

        keepalive_timeout  65;

        #gzip  on;

        upstream monitors {
            server monitoring-system-service.kube-apps:8080;
        }

        server {
            listen 80;
            auth_basic           "closed site";
            auth_basic_user_file ngauth;
            location / {
              if ($request_method = OPTIONS ) {
                add_header Access-Control-Allow-Origin "null"; # <- needs to be updated
                add_header Access-Control-Allow-Methods "GET, OPTIONS";
                add_header Access-Control-Allow-Headers "Authorization";   # <- You may not need this...it's for Basic Auth
                add_header Access-Control-Allow-Credentials "true";        # <- Basic Auth stuff, again
                add_header Content-Length 0;
                add_header Content-Type text/plain;
                return 200;
              }
              proxy_pass http://monitors;
            }
        }
    }
  ngauth: |
    username:password
```
*注意:ngauth 下面的username和password需要替换为步骤1中生成文件的用户名和密码*
3. 建议在本地先使用docker进行调试，如果没有问题，则可以进行下一步，部署到k8s中
```
docker run --name nginx-container -v /home/user/nginx/:/etc/nginx/nginx.conf:ro -d nginx:1.12.2

```
4. 部署到k8s
nginx.yaml如下：
    ```
        apiVersion: v1
        kind: Service
        metadata:
          name: nginx
          labels:
            app: nginx
          namespace: kube-apps
        spec:
          type: NodePort
          selector:
            app: nginx
          ports:
          - name: http
            port: 80
            targetPort: 80
        \---
        apiVersion: extensions/v1beta1
        kind: Deployment
        metadata:
          name: nginx
          namespace: kube-apps
          labels:
            addonmanager.kubernetes.io/mode: Reconcile
        spec:
          template:
            metadata:
              labels:
                app: nginx
            spec:
              containers:
              - name: nginx
                image: nginx:1.12.2
                ports:
                - containerPort: 80
                volumeMounts:
                - name: config-volume
                  mountPath: /etc/nginx/
              volumes:
              - name: config-volume
                configMap:
                  name: nginx-conf
                  items:
                  - key: nginx.conf
                    path: nginx.conf
                  - key: ngauth
                    path: ngauth
    ```
执行：
```
$ kubectl create -f configMap.yaml 
$ kubectl create -f nginx.yaml 
```
5. 配置traefik指向地址为nginx的service地址即可
### 参考
[http://nginx.org/en/docs/http/ngx_http_auth_basic_module.html](http://nginx.org/en/docs/http/ngx_http_auth_basic_module.html)