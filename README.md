# liongo
![-_-](https://travis-ci.org/Anteoy/liongo.svg?branch=master)

liongo,a engine that can power for your blog by [Golang](https://golang.org).

**Current: 1.2.0 (Beta) 2017.03.08**

Liongo now supports the following features

##build
```
liongo build
```
that can generate a simple site static file
##run
```
liongo run
```
It's contains the "liongo build",it will use the build file to run a serve at default port 8080,you can also use
```
##run --note
```
liongo run --note
```
It's can open you online note,you must install the mongodb and mysql in your server,before this,you should see the config of the code.
liongo run :[your port]
```
set the run port
##new
```
liongo new [yourblogtitle]
```
It will generate the article file in the build file,and you can update the file to build you site
