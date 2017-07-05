---
date: 2016-09-09 23:23:00
title: maven手动安装远程仓库中找不到的jar到本地仓库
categories:
    - 架构
tags:
---

只需要使用一个maven命令即可：
 mvn install:install-file -Dfile=（jar本地地址）/root/IdeaProjects/keta-customer/lib/patchca-0.5.0.jar -DgroupId=com.github.bingoohuang（groupId） -DartifactId=(artifactId)patchca -Dversion=(version)0.5.0 -Dpackaging=jar
即可利用pom像普通情况一样使用，在项目里自由使用了

后来因需安装bairong jar
使用命令：
mvn install:install-file -Dfile=/Users/daniel/Desktop/bsApi-2.3.5-shaded.jar -DgroupId=bairongapi -DartifactId=bsApi -Dversion=2.3.5 -Dpackaging=jar -DgeneratePom=true
mvn install:install-file -Dfile=/Users/daniel/Desktop/bsApi-2.4.0-shaded.jar -DgroupId=bairongapi -DartifactId=bsApi_test -Dversion=2.4.0 -Dpackaging=jar -DgeneratePom=true
mvn install:install-file -Dfile=//Users/cool/work/lib/jaxb-api-2.2.jar -DgroupId=jaxb-api -DartifactId=javax-jaxb-api -Dversion=2.2 -Dpackaging=jar -DgeneratePom=true
期间报错:
[ERROR] Failed to execute goal org.apache.maven.plugins:maven-install-plugin:2.4:install-file (default-cli) on project standalone-pom: Error installing artifact 'bairongapi:bsApi_test:jar': Failed to install artifact bairongapi:bsApi_test:jar:2.4.0: /home/zhoudazhuang/repository/bairongapi/bsApi_test/2.4.0/bsApi_test-2.4.0.jar (权限不够) -> [Help 1]
[ERROR]
[ERROR] To see the full stack trace of the errors, re-run Maven with the -e switch.
[ERROR] Re-run Maven using the -X switch to enable full debug logging.
[ERROR]
[ERROR] For more information about the errors and possible solutions, please read the following articles:
[ERROR] [Help 1] http://cwiki.apache.org/confluence/display/MAVEN/MojoExecutionException
于是使用-e参数：
mvn install:install-file -e -Dfile=/home/zhoudazhuang/zhaolei/bsApi-2.4.0-shaded.jar -DgroupId=bairongapi -DartifactId=bsApi_test -Dversion=2.4.0 -Dpackaging=jar -DgeneratePom=true
得到提示：
[ERROR] Failed to execute goal org.apache.maven.plugins:maven-install-plugin:2.4:install-file (default-cli) on project standalone-pom: Error installing artifact 'bairongapi:bsApi_test:jar': Failed to install artifact bairongapi:bsApi_test:jar:2.4.0: /home/zhoudazhuang/repository/bairongapi/bsApi_test/2.4.0/bsApi_test-2.4.0.jar (权限不够) -> [Help 1]
org.apache.maven.lifecycle.LifecycleExecutionException: Failed to execute goal org.apache.maven.plugins:maven-install-plugin:2.4:install-file (default-cli) on project standalone-pom: Error installing artifact 'bairongapi:bsApi_test:jar': Failed to install artifact bairongapi:bsApi_test:jar:2.4.0: /home/zhoudazhuang/repository/bairongapi/bsApi_test/2.4.0/bsApi_test-2.4.0.jar (权限不够)
于是切至root，成功install

