---
date: 2016-04-10 10:19:00
title: a bit of git
categories:
    - 架构
tags:
    - git
---

#a bit of git

*nano 编辑器 git pull = git fetch + git merge*

1. *关于 .gitignore*
规则很简单，不做过多解释，但是有时候在项目开发过程中，突然心血来潮想把某些目录或文件加入忽略规则，若更过后未生效，原因是.gitignore只能忽略那些原来没有被track的文件，如果某些文件已经被纳入了版本管理中，则修改.gitignore是无效的。那么解决方法就是先把本地缓存删除（改变成未track状态），然后再提交



 1.  git rm -r --cached .
 2. git add .
 3. git commit -m 'xx'
2. 解决冲突方法一(利用stash)
 1. git stash list
 2. git stash save [filename]
 3. git stash pop [list-id]
 4. *然后解决冲突并提交*
3. 解决冲突方法一(利用commit)


```
root@zhoudazhuang-pc:/root/IdeaProjects/jottings# git pull

remote: Counting objects: 18, done.

remote: Compressing objects: 100% (5/5), done.

remote: Total 18 (delta 8), reused 17 (delta 7), pack-reused 0

展开对象中: 100% (18/18), 完成.

来自 https://github.com/Anteoy/jottings

   4b57c65..615069d  master     -> origin/master

更新 4b57c65..615069d

error: Your local changes to the following files would be overwritten by merge:

        src/main/java/com/anteoy/sample/Static.java

Please, commit your changes or stash them before you can merge.

Aborting

root@zhoudazhuang-pc:/root/IdeaProjects/jottings# git add .

root@zhoudazhuang-pc:/root/IdeaProjects/jottings# git commit -m 'dsfsf'

[master f37b17c] dsfsf

 1 file changed, 1 insertion(+), 1 deletion(-)

root@zhoudazhuang-pc:/root/IdeaProjects/jottings# git pull

自动合并 src/main/java/com/anteoy/sample/Static.java

冲突（内容）：合并冲突于 src/main/java/com/anteoy/sample/Static.java

自动合并失败，修正冲突然后提交修正的结果。

同上一样解决冲突并提交
```
4.git rebase
 用于回溯版本控制或进行分支合并