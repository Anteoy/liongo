---
date: 2017-03-25 00:40:00
title: 简析超文本传输协议（HTTP，HyperText Transfer Protocol)
categories:
    - 协议
tags:
    - HTTP
---

## http
超文本传输协议（HTTP，HyperText Transfer Protocol)
## http request结构：
1. request line: 请求行 包含请求的方法（如get,post） 请求资源路径（URL,URL总是以/开头，/就表示首页） HTTP协议版本号
2. request head: 其他重要请求信息 如服务器生成的response给浏览器的cookie，后面的请求携带在request head中（Cookie是由服务器创建的，然后通过response响应发送给客户端的一个键值对，session在服务端）
    ```
        authority:www.google.co.jp
        :method:GET
        :path:/search?q=cookie+session%E5%9C%A8body%E8%BF%98%E6%98%AFhead%E4%BC%A0%E8%BE%93&oq=cookie+session%E5%9C%A8body%E8%BF%98%E6%98%AFhead%E4%BC%A0%E8%BE%93&aqs=chrome..69i57.223j0j4&sourceid=chrome&ie=UTF-8
        :scheme:https
        accept:text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
        accept-encoding:gzip, deflate, sdch, br
        accept-language:zh-CN,zh;q=0.8
        avail-dictionary:3GC5dhWe
        cookie:SID=fASEd608GMFg0RBNPi3M2kcukdv65QmfliLImKmzpsMjAWlOTOP1CDyp_N8S1nzogLU-Cg.; HSID=AisZXXax41iB7UNy4; SSID=Adil26yAHJqVzMcwV; APISID=OVxpy9BlM75Y8U6d/AT6FZmCUVzKW01zjw; SAPISID=9sgr8Ja05vRr8X41/AXQPIDREewmR2Tara; NID=99=YPixFcm5pI3MK1q4u1tRuNqd0SyCi0l504phONJs9ZyzPGiETkZ-by5wVaHwc6-D63O1ucvCaytqIgrVrW84TKJ20HoGJ6zdeqHURUk_2e85TKNRPoZp6sKi1tgKoA1bs1FYduCLKN1j3p1UyOFPX8n6c_NgLEKOaZgqcqdfWZXPRHu30AYt1HImlcY7pBpWRJUDUWF8p0LcIC4o
        upgrade-insecure-requests:1
        user-agent:Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
        x-client-data:CIS2yQEIprbJAQiCmcoBCKmdygE=
    ```
3. 携带提交给web服务器的数据。使用GET方法时，为空。Body和Header之间空一行。
## http response结构：
1. response line: HTTP/version-number：HTTP协议版本号;status-code：状态码，反应服务器处理是否正常，告知出现的错误;message：状态消息，同状态码对应。
2. response head:其他重要响应信息 如
    ```
        Accept-Ranges:bytes
        Access-Control-Allow-Origin:*
        Age:1987
        Connection:keep-alive
        Content-Encoding:gzip
        Content-Length:5310
        Content-Security-Policy:default-src 'none'; style-src 'unsafe-inline'; img-src data:; connect-src 'self'
        Content-Type:text/html; charset=utf-8
        Date:Fri, 24 Mar 2017 15:54:11 GMT
        ETag:W/"5822212c-247c"
        Server:GitHub.com
        Vary:Accept-Encoding
        Via:1.1 varnish
        X-Cache:HIT
        X-Cache-Hits:1
        X-Fastly-Request-ID:8427de5ab5a90c8efdfea6401aeb4bf8765ba256
        X-GitHub-Request-Id:BD16:06EC:DB5E4C5:115647E3:58D5395F
        X-Served-By:cache-nrt6127-NRT
        X-Timer:S1490370851.712212,VS0,VE0
    ```
    KeepAlive 属性可以有效地降低TCP握手的次数
3. response body: 包含响应的内容，网页的HTML源码就在Body中。Body的数据类型由Content-Type头来确定，如果是网页，Body就是文本，如果是图片，Body就是图片的二进制数据。
当存在Content-Encoding时，Body数据是被压缩的，最常见的压缩方式是gzip，所以，看到Content-Encoding: gzip时，需要将Body数据先解压缩，才能得到真正的数据。压缩的目的在于减少Body的大小，加快网络传输。

## postman post请求中body的数据类型
1. mutipart/form-data
    
    网页表单用来传输数据的默认格式。可以模拟填写表单，并且提交表单。
    可以上传一个文件作为key的value提交(如上传文件)。但该文件不会作为历史保存，只能在每次需要发送请求的时候，重新添加文件。
2. urlencoded

    同前面一样，注意,你不能上传文件通过这个编码模式。
    该模式和表单模式会容易混淆。urlencoded中的key-value会写入URL，form-data模式的key-value不明显写入URL，而是直接提交。
3. raw

    raw request可以包含任何东西。所有填写的text都会随着请求发送。
4. binary

    image, audio or video files.text files 。 也不能保存历史，每次选择文件，提交。
## 后记
postman post请求中body的数据类型参考自 [Postman 详解](http://www.jianshu.com/p/35678284ce78)n 详解](http://www.jianshu.com/p/35678284ce78)