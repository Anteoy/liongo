---
date: 2017-10-04 22:49:00
title: opensuse下安装网易云音乐netease-cloud-music
categories:
    - linux
tags:
    - linux,archlinux,网易云音乐
---

### 安装过程
1. 安装环境
opensuse 42.3（理论上仓库中列出的支持版本都可以用此方法安装）
2. 官方软件仓库搜索netease，也可直接点击此处打开：https://software.opensuse.org/package/netease-cloud-music?search_term=netease
3. 选择对应版本，点击Source下载源码，我这里下载的对应版本为42.3，得到如下文件（注意这里点击1 Click Install安装会出错，我最开始点击直接安装，结果提示安装成功，但无法使用，仔细看安装日志会发现，虽然最后提示安装成功了，但其实在过程中已经报错）
    ```
    -rwxrwxrwx 1 zhoudazhuang users     7516 10月  4 21:37 netease-cloud-music-1.0.0-8.7.src.rpm
    ```
4. 解压得到以下两个文件，其中一个就是安装脚本netease-cloud-music.sh.in：
    ```
    -rwxrwxrwx 1 zhoudazhuang users     1975 3月   7 2017 netease-cloud-music.sh.in
    -rwxrwxrwx 1 zhoudazhuang users    11840 7月  11 06:42 netease-cloud-music.spec
    ```
5. 脚本wget无法直接下载需要的安装deb包，所以手动在浏览器下载对应文件，其中版本必须与脚本中需要安装的版本一致，即1.0.0，点击如下连接下载：
    ```
    http://s1.music.126.net/download/pc/netease-cloud-music_1.0.0_amd64_ubuntu16.04.deb
    ```
6. 把下载的文件放入/tmp目录下
7. 执行脚本安装
    ```
    sudo ./netease-cloud-music.sh.in
    ```
    输出：
    ```
    [sudo] root 的密码：
    Downloading deb package from netease ...
    wget: unrecognized option '--show-progress'
    用法： wget [选项]... [URL]...
    
    请尝试使用“wget --help”查看更多的选项。
    Successfully downloaded  to /tmp/netease-cloud-music_1.0.0_amd64_ubuntu16.04.deb.
    Unpacking netease-cloud-music_1.0.0_amd64_ubuntu16.04.deb ... it'll take some time ...
    Successfully unpacked /tmp/netease-cloud-music_1.0.0_amd64_ubuntu16.04.deb to /tmp/netease-cloud-music-1.0.0/usr
    Congratulations! Installation succeed!
    ```

### 参考：[opensuse netease网易云音乐](https://www.douban.com/note/594581641/)