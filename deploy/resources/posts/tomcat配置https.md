---
date: 2016-11-21 17:47:00
title: tomcat配置https
categories:
    - 架构
tags:
---

﻿**搭建环境：**


 1. ubuntu 16.04 LTS

 2. apache tomcat 7

 3. java 7

**搭建过程：**
1. 服务端利用jdk自带的keytool生成server.keystore,命令如下：
	`keytool -genkey -alias tomcat -keyalg RSA -keypass anteoypasswd -storepass Envisi0n -keystore server.keystore -validity 3600`
	部分参数说明如下
	- alias <alias>                  要处理的条目的别名
	- keyalg <keyalg>                密钥算法名称
	- keysize <keysize>              密钥位大小
	- sigalg <sigalg>                签名算法名称
	- destalias <destalias>          目标别名
	- dname <dname>                  唯一判别名
	- startdate <startdate>          证书有效期开始日期/时间
	- ext <value>                    X.509 扩展
	- validity <valDays>             有效天数
	- keypass <arg>                  密钥口令
	- keystore <keystore>            密钥库名称
	- storepass <arg>                密钥库口令
	- storetype <storetype>          密钥库类型
	- providername <providername>    提供方名称
	- providerclass <providerclass>  提供方类名
	- providerarg <arg>              提供方参数
	- providerpath <pathlist>        提供方类路径
	- v                              详细输出
	- protected                      通过受保护的机制的口令
	        注：期间填写名字组织国家地区，本人是随意填写的，最后Y确认，测试通过，上面密码后面tomcat配置需要使用到。
2. 使用server.keystore生成.cer无私钥证书（pfx为有公私钥的证书）

    `keytool -export -trustcacerts -alias tomcat -file server.cer -keystore server.keystore -storepass anteoypasswd`
3. 配置tomcat

    conf\server.xml

    找到

    `<Connectorport="8443" protocol="HTTP/1.1"SSLEnabled="true" maxThreads="150" scheme="https" secure="true" clientAuth="false" sslProtocol="TLS" />`

   取消注释并增加(文件位置以及之前使用的keystore密码)：

    `keystoreFile="/home/zhoudazhuang/test/apache-tomcat-7.0.72/server.keystore" keystorePass="anteoypasswd"`

    最终如下：

    `<Connector port="8443" protocol="org.apache.coyote.http11.Http11Protocol"

               maxThreads="150" SSLEnabled="true" scheme="https" secure="true"

               clientAuth="false" sslProtocol="TLS" keystoreFile="/home/zhoudazhuang/test/apache-tomcat-7.0.72/server.keystore" keystorePass="anteoypasswd"		/>`
4. 将证书导入到JDK的cacerts库

    `sudo keytool -import -trustcacerts -alias tomcat -file server.cer -keystore /usr/lib/jvm/java-7-openjdk-amd64/jre/lib/security/cacerts -storepass changeit`
5. 和普通方式一样在webapp中部署web项目，使用https://localhost:8443/+项目，信任此网站即可正常使用https访问


