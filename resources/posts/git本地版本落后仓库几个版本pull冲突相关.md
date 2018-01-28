---
date: 2016-07-19 17:24:00
title: git本地版本落后仓库几个版本pull冲突相关
categories:
    - 架构
tags:
    - git
---

git中本地落后仓库几个版本pull冲突,但是自己又不想提交本地的代码到远程仓库，可以尝试使用下面的方法：
一种是使用git fetch，但是自己这个用得少
另外就是使用git pull
git pull = git fetch + git merch
解决冲突时利用git stash 把本地代码保存起来
1.git pull origin master
报错：有冲突 不能拉取下来（本地和远程在同地方都有改动 ）

```
来自 https://git.coding.net/zhoudafu/ISM_D
 * branch            master     -> FETCH_HEAD
更新 ac5fccd..170cf6f
error: The following untracked working tree files would be overwritten by merge:
	.idea/artifacts/ISM_war_exploded.xml
Please move or remove them before you can merge.

```
2. git stash
3. git pull origin master //这个时候一直报下面错误

```
更新 ac5fccd..170cf6f
error: The following untracked working tree files would be overwritten by merge:
	.idea/artifacts/ISM_war_exploded.xml
Please move or remove them before you can merge.
Aborting

```
说明ISM_war_exploded.xml 这个文件没有stash进去
4.使用 git stash -a

```
root@zhoudazhuang-PC:~/IdeaProjects/ISM_D# git stash -a
Saved working directory and index state WIP on master: ac5fccd beta1.0
HEAD 现在位于 ac5fccd beta1.0
root@zhoudazhuang-PC:~/IdeaProjects/ISM_D# git pull
更新 ac5fccd..170cf6f
Fast-forward
 .idea/artifacts/ISM_war_exploded.xml             |  59 +++++++++++++
 .idea/libraries/library_app.xml                  |  10 +++
 .idea/libraries/mail_app.xml                     |  13 +++
 ImageWeb.war                                     | Bin 0 -> 19882832 bytes
 ImageWeb_war.war                                 | Bin 0 -> 19882832 bytes
 WebContent/main/js/chat.js                       |   1 +
 WebContent/main/main.html                        | 104 +++++++++++------------
 src/com/zy/web/ism/mapper/xml/EmployeeMapper.xml |   2 +-
 8 files changed, 136 insertions(+), 53 deletions(-)
 create mode 100644 .idea/artifacts/ISM_war_exploded.xml
 create mode 100644 .idea/libraries/library_app.xml
 create mode 100644 .idea/libraries/mail_app.xml
 create mode 100644 ImageWeb.war
 create mode 100644 ImageWeb_war.war
root@zhoudazhuang-PC:~/IdeaProjects/ISM_D# git stash list
stash@{0}: WIP on master: ac5fccd beta1.0

```
成功
5.使用 git pull origin master 把主干拉下来

6.使用git stash list
  找到你刚刚stash的id 比如我的是


```
root@zhoudazhuang-PC:~/IdeaProjects/ISM_D# git stash list
stash@{0}: WIP on master: ac5fccd beta1.0

```
7.使用git stash pop stash@{0} 取出刚刚存入代码 如果无冲突将会自动合并 如果有冲突需要你进入文件手动解决冲突，冲突在冲突文件夹里会有明显标注你本地代码和仓库代码


