---
date: 2018-01-09 17:14:00
title: kubernetes下用logstash代替efk中的fluentd
categories:
    - kubernetes
tags:
    - kubernetes,nginx,k8s,traefik
---

### 前言
目前我们的系统架构，从阿里云的docker compose迁移到了自建的kubernetes,而我们的日志系统也几经周折。从最开始的阿里云docker compose环境下的elk + kafka，使用了一段时间后由于老板觉得阿里云的kafka收费价格不怎么划算，并且线上服务器的资源吃紧，然后切换到了阿里云的日志服务。线上切换到kubernetes以后，又使用efk搜集了一段时间日志，后来发现fluentd搜集的日志存在一些延迟问题和准确性问题，并且变更配置搜集特定的日志时较为繁琐，于是准备把fluentd换为logstash + rabbit,期间对比了下redis,rabbit和kafka，最终选择了在用的自建的rabbit集群，由于近期事情较多，精力有限，此文只会阐述关于logstash和rabbit相关使用的过程，关于rabbit自建集群，以及k8s的efk，以及elk搭建，后面会逐一补充。

#### 简要步骤
1. 编写logstash的配置文件logstash.conf 建议先在本地使用docker部署logstash，然后在本地进行测试，成功后再部署到kubernetes环境
    ```
        input {
            rabbitmq {
              type => "app-log"
              host => "your rabbitmq host"
              user => "your user"
              password => "your password"
              port => 5672
              exchange => "exchange_logs"
              exchange_type => "direct"
              queue => "queue_logs"
              ack => false
              durable => true
            }
           }
        
           output {
            elasticsearch {
              hosts => ["elasticsearch.kube-ops:9200"]
              index => "logstash-%{type}-%{+YYYY.MM.dd}"
              document_type => "%{type}"
              flush_size => 20000
              idle_flush_time => 10
              template_overwrite => true
            }
            stdout {
                codec => rubydebug
            }
           }
    ```
output同时输出到es和控制台，便于调试
2. 编写logstash的k8s configMap:
    ```
        kind: ConfigMap
        apiVersion: v1
        metadata:
          name: logstash-conf
          namespace: kube-ops
        data:
          logstash.conf: |
           input {
            rabbitmq {
              type => "app-log"
              host => "rabbitmq...."
              user => "..."
              password => "..."
              port => 5672
              exchange => "exchange_logs"
              exchange_type => "direct"
              queue => "queue_logs"
              ack => false
              durable => true
            }
           }
        
           output {
            elasticsearch {
              hosts => ["elasticsearch.kube-ops:9200"]
              index => "logstash-%{type}-%{+YYYY.MM.dd}"
              document_type => "%{type}"
              flush_size => 20000
              idle_flush_time => 10
              template_overwrite => true
            }
            stdout {
                codec => rubydebug
            }
           }
    ```
    执行
    ```
        kubectl create -f logstash-configMap.yaml
    ```
3. 建议在本地先使用docker进行调试，如果没有问题，则可以进行下一步，部署到k8s中
    ```
        docker run -d  -v ~/a/docker/efk-deploy/1.7/bac/es2/qatrans/logstash/logstash.conf:/etc/logstash/logstash.conf  logstash:5.6.4 logstash -f /e
        tc/logstash/logstash.conf
    
    ```
4. 部署到k8s
    logstash.yaml如下：
    ```
        apiVersion: v1
        kind: Service
        metadata:
          name: logstash
          namespace: kube-ops
        spec:
          ports:
          - port: 5044
            targetPort: beats
          selector:
            type: logstash
          clusterIP: None
        \---
        apiVersion: extensions/v1beta1
        kind: Deployment
        metadata:
          name: logstash
          namespace: kube-ops
        spec:
          template:
            metadata:
              labels:
                type: logstash
            spec:
              containers:
              - image: anteoy/logstash:5.6.4
                name: logstash
                ports:
                - containerPort: 5044
                  name: beats
                command:
                - logstash
                - '-f'
                - '/etc/logstash_c/logstash.conf'
                volumeMounts:
                - name: config-volume
                  mountPath: /etc/logstash_c/
                resources:
                  limits:
                    cpu: 1000m
                    memory: 2048Mi
                  # requests:
                  #   cpu: 512m
                  #   memory: 512Mi
              volumes:
              - name: config-volume
                configMap:
                  name: logstash-conf
                  items:
                  - key: logstash.conf
                    path: logstash.conf
    ```
    执行：
        ```
        $ kubectl create -f logstash.yaml
        ```
### 参考
[https://www.elastic.co/guide/en/logstash/current/plugins-inputs-rabbitmq.html](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-rabbitmq.html)
[https://gist.github.com/dblessing/a9d5a68da56eb451553a](https://gist.github.com/dblessing/a9d5a68da56eb451553a)