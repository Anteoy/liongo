---
date: 2017-06-25 23:48:00
title: git submodule以及被忽略的.gitignore
categories:
    - golang
tags:
    - golang,new,make
---

###引言：
　不积跬步，无以至千里；不积小流，无以成江海。原本自以为对git相关的东西以及原理已经有了不错的认识和理解，然而当碰到一个小小的problem才发现自己掌握得并没有自己想象的那么好，好记性不如烂笔头，于是决定把自己对.gitignore以及submodule的部分理解简单记录下来，由于个人能力有限，如有不当之处，欢迎指正。
　
###git submodule

git submodule在项目中虽一直在使用（用于维护自有公共类库），不过对此知之甚少，于是抽空做一个小小的总结。可运行git submodule --help获取帮助doc。
目地：多项目维护共用公共类库starjazz
步骤：
####在项目中初始化submodule
1. 在即将引用的项目中初始化此git submodule
	git submodule add http://gitlab-ce.yougitlabname.com/backend/starjazz
	git status 
	```
	新文件：   .gitmodules
	新文件：   starjazz
	```
	 生成的.gitmodules文件中记录了submodules的path路径以及url指针,同时在.git文件夹下你能找到你刚刚添加的子模块starjazz
 2. cd starjazz && git status 
	 ```
	 mkdir: 已创建目录 '/home/zhoudazhuang/.zsh_history/home/zhoudazhuang/x/wechatMetric/starjazz'
位于分支 master
您的分支与上游分支 'origin/master' 一致。
无文件要提交，干净的工作区
	 ```
	 即可作为一个独立的公共类库进行提交和变更。
3. 在主（父）项目中提交加入的git submodule
	- cd .. && git status
	- git add .
	- git commit -m 'submodule'
	- git push origin test
	提交后会在主（父）项目中生成对submodule的引用，但不会提交submodule内部代码。
####pull 带有submodule的项目
1. git clone http://gitlab-ce.youname.com/backend/wechatMetric
	此时submodule starjazz为一个没有任何文件的空文件夹
2. 执行如下拉取子模块代码
	```
	init 会把.gitsubmodule中的配置注册到.git/config文件中
	$ git submodule init  
	子模组 'starjazz' (http://gitlab-ce.youname.com/backend/starjazz) 未
	对路径 'starjazz' 注册
	
	# zhoudazhuang at zhoudazhuang-pc in ~/youname/submodule-test/wechat
	Metric on git:test o [22:58:32]
	$ git submodule update
	正克隆到 'starjazz'...
	Username for 'http://gitlab-ce.youname.com': anteoy@youname.com
	Password for 'http://anteoy@youname.com@gitlab-ce.youname.com':
	remote: Counting objects: 385, done.
	remote: Compressing objects: 100% (331/331), done.
	remote: Total 385 (delta 78), reused 213 (delta 6)
	接收对象中: 100% (385/385), 252.75 KiB | 0 bytes/s, 完成.
	处理 delta 中: 100% (78/78), 完成.
	检查连接... 完成。
	子模组路径 'starjazz'：检出 '21d99fdd97b76bbb524c264aa619431143b6cf83
	'
	```
3. cd starjazz && git status
	```
	mkdir: 已创建目录 '/home/zhoudazhuang/.zsh_history/home/zhoudazhuang/
	youname/submodule-test/wechatMetric/starjazz'
	头指针分离于 21d99fd
	无文件要提交，干净的工作区
	```
	注意：git submodule update检出项目的指定版本（HEAD），并不指向一个分支。头指针和分支未绑定，是分离状态。（并且每次执行submodule update都会从git仓库覆盖本地已变更的代码），解决如下
	- 强制将 master 分支指向当前头指针的位置
		```
		# 强制将 master 分支指向当前头指针的位置
		$ git branch -f master HEAD
		# 检出 master 分支
		$ git checkout master
		```
		即可正常进行代码提交.
4. git submodule add http://gitlab-ce.youname.com/backend/starjazz src/vendor/testsub/starjazz 可设置子模块在当前项目的相对路径.
5. submodule有更新，在其他项目拉取代码操作（同时也可以进入submodule项目进行处理）
	 foreach为每一个递归循环更新recursive所有的submodule(这里仅执行git submodule update是无效的，此update只在第一次拉去更新.git/config文件的子模块有效)
	```
	git submodule foreach --recursive git pull origin master
	```
6. git clone --recursive ... 可在拉取项目时把依赖的子模块同时拉取下来.
7. 更新.gitsubmodule中的remote url使用git submodule sync。
8. 删除submodule需要使用git rm --cached删除仓库文件,rm删除本地文件，删除.gitsubmodule文件在红对应的子项目，删除.git/config文件下的git配置文件。(注意:如果删除文件后未执行git add操作,则重新加入git submodule会报错'src/vendor/gitlab-ce.youname.com/backend/starjazz' 已经存在于索引中,此时执行git add ,commit即可)

### 被忽略的.gitignore：

 由一个文件引起的文件夹代码未提交
	.gitignore文件增加：
	```
	test1
	bac
	customerservice
	```
	对于customerservice，git仓库中存在一个文件和文件夹（原目地为忽略项目中生成的customerservice），最后此命名的文件和文件夹都被git忽略。期间尝试使用/customerservice，customerservice/，\*/customerservice/\*未果，最后参阅gitignore官方文档，地址：[https://git-scm.com/docs/gitignore](https://git-scm.com/docs/gitignore),/customerservice，\*/customerservice/\*这两种为无效用法，/customerservice/这种也无效,,\*\*/customerservice 此为忽略文件和文件夹，\**/customerservice/\**，customerservice/ ,**/customerservice/这几种效果一样,也是同时忽略文件和文件夹。
### 小结及延伸：
1. 忽略某文件 *.exe,src/wechatmetric/microservice/customerservice/customerservice(从.git文件夹同目录开始)
2. 同时忽略文件夹和文件 \**/customerservice/\**，customerservice/,\*\*/customerservice , customerservice（不包含斜杠的命名，如3中customerservice，则按shell查找命令同时忽略）
4. 将匹配过的部分移出 ！操作符 ，如 忽略aa文件夹，但排除其中bb：
	```aa/
	!aa/bb/
	```
