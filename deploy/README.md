## 当前测试版服务器发布说明
1. 打包
    `export GOPATH=$PWD`
    `go build -o blog_server github.com/Anteoy/liongo/main/`
2. 按照示例deploy拷贝资源文件到deploy目录
3. 将deploy目录压缩后发送至服务器
    `scp ./deploy.tar.gz name@ip:/home/...`
4. 解压后进入目录运行
    `./blog_server run --note`
5. 使用服务器ip:8080即可访问.
## 先决条件：
1. 有可用的mongodb和mysql，配置好后再进行go build。
2. mongodb需要有测试库liongo,集合note以及一个测试sample文档，内容如下：
    `{
    "_id" : ObjectId("5883232b7b71193a4e742442"),
    "name" : "test1",
    "content" : "<p>#For my memory scan ##1. 右上角资源监测 避免公司低配资源占用问题</p>\n"
    }`
3. 服务器防火墙请自行开放8080端口，或者使用nginx等代理服务器作转发.