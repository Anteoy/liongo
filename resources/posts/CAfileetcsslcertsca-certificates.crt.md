---
date: 2016-05-11 9:19:42
title: CAfileetcsslcertsca-certificates.crt
categories:
    - 架构
tags:
    - problem
---

#问题描述

- CAfile:/etc/ssl/certs/ca-certificates.crt CRLfile:none

- ubuntu 16.04

- git https

#解决方案

 1. export GIT_SSL_NO_VERIFY = 1

 2. sudo update-ca-certificates

 3. echo -n | openssl s_client -showcerts -connect github.com:443 2>/dev/null | sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p'

    1. echo -n  *n表示输出文字后不换行*

    2. 管道符仅处理前一指令输出正确的信息，对standard error无法直接处理，后交接至下一指令

    3. 2>/dev/null 把错误信息存入黑洞，直接抛弃，正确的才会通过下一个管道符进入下一指令

    4. &表示后台执行，你可以继续占有你的输入窗口

    5. sed 是一种在线编辑器 一次一行，当前行存入缓冲区（称为模拟空间）再用sed指令处理缓冲区内容

        - -e 允许多台编辑

        - -n 不打印，不写编辑行到标准输出

    6. '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' 设置打印范围

#致谢

[stackoverflow](http://stackoverflow.com/questions/21181231/server-certificate-verification-failed-cafile-etc-ssl-certs-ca-certificates-c)




