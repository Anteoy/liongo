---
date: 2016-01-27 14:58:00
title: git branch
categories:
    - 架构
tags:
    - 架构
---

###从git拉取指定分支
1. 先clone下来
	```git clone https://git.coding.net/zhoudafu/gblog.git```
2. 增加远程仓库（此步可省略，若添加远层仓库可参考）
	```git remote add origin https://git.coding.net/zhoudafu/gblog.git```
3. fetch下分支代码
	```git fetch origin coding-pages```
4. 使用checkout切换分支
	```git checkout -b coding-pages origin/coding-pages```
	```分支 gh-pages 设置为跟踪来自 origin 的远程分支 coding-pages。切换到一个新分支 'coding-pages'```
	注：创建分支： git branch mybranch
			切换分支： git checkout mybranch
			创建并切换分支： git checkout -b mybranch
5. 使用git status,add,commit,push origin gh-pages正常提交

###本地创建分支并提交到远程分支
1. 从已有的分支创建新的分支(如从master分支),创建一个gh-pages分支，执行完成自动切换到gh-pages分支，可使用git branch查看和切换
	```git checkout -b coding-pages```
2. 提交到远程仓库
	```git push origin coding-pages```
###添加多远程仓库 并拉取，合并，提交其他远程仓库代码

```
 1823  git remote add origin2 https://github.com/yan-chou-strong/blog.git
 1824  git fetch origin2 master
 1825  git checkout -b master2 origin2/master
 1826  git status
 1827  git merge master
 1828  vim README.md
 1829  git status
 1830  git add README.md
 1831  git status
 1832  git add .
 1833  git commit -m 'another T'
 1834  git push origin2 master
 1835  git pull
 1836  git push origin2 master
 1837  git pull
 1838  git push origin2 master
 1839  git add .
 1840  git push origin2 master
 1841  git pull origin2 master
 1842  git status
 1843  git push origin2 master
 1844  git push origin master
 1845  git status
 1846  git pull origin2/master
 1847  git pull origin2 master
 1848  git fetch origin2 master
 1849  git merge origin2/master
 1850  git push origin2 master
 1851  git status
 1852  git push origin2 master2

```

###分支merge
1. git checkout master
2. git merge coding-pages
###同步fork项目
1. git remote add username https://github.com/xxx.git
2. git fetch username
3. git merge username/master
###后记
*在我使用pugo push时 ，使用本地新建分支再提交到远程*
```EROR|12-27 20:52:07.7178|Git|Fail|warning: push.default 尚未设置，它的默认值在 Git 2.0 已从 'matching'
变更为 'simple'。若要不再显示本信息并保持传统习惯，进行如下设置：
  git config --global push.default matching
若要不再显示本信息并从现在开始采用新的使用习惯，设置：
  git config --global push.default simple
当 push.default 设置为 'matching' 后，git 将推送和远程同名的所有
本地分支。
从 Git 2.0 开始，Git 默认采用更为保守的 'simple' 模式，只推送当前
分支到远程关联的同名分支，即 'git push' 推送当前分支。
参见 'git help config' 并查找 'push.default' 以获取更多信息。
（'simple' 模式由 Git 1.7.11 版本引入。如果您有时要使用老版本的 Git，
为保持兼容，请用 'current' 代替 'simple'）
fatal: 当前分支 coding-pages 没有对应的上游分支。
为推送当前分支并建立与远程上游的跟踪，使用
   git push --set-upstream origin coding-pages
``
*后在gitPro目录下使用*
``` git push --set-upstream origin coding-pages```
回执
```
Delta compression using up to 4 threads.压缩对象中: 100% (23/23), 完成.写入对象中: 100% (29/29), 343.86KiB | 0 bytes/s, 完成.Total 29 (delta 5), reused 0 (delta 0)To https://git.coding.net/zhoudafu/gblog.git 6363c13..5f88c38  coding-pages -> coding-pages
分支 coding-pages 设置为跟踪来自 origin 的远程分支 coding-pages
```
重新执行未果报错
```
EROR|12-27 20:57:08.0745|Git|Fail|exit status 1
```
在这里应该是pugo内部调用处理有差异，使用第一种分支操作即可成功提交。

