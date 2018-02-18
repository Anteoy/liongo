# liongo
![-_-](https://travis-ci.org/Anteoy/liongo.svg?branch=master)

liongo,an engine that can power for your blog by [Golang](https://golang.org).

**Current: 1.2.0 (Beta) 2017.03.08**

## how to get liongo

1. setup your gopath and gobin
2. go get -x github.com/Anteoy/liongo
3. go install  github.com/Anteoy/liongo/liongo
4. copy the folder "resources" to the parent directory of folder gobin
## how to use liongo

Liongo now supports the following features
### build
```
liongo build
```
that can generate a simple site static file
### run
```
liongo run
```
It's contains the "liongo build",it will use the build file to run a serve at default port 8080,you can also use
```
### run --note(need your server's mysql and mongodb supported)
```
liongo run --note
```
It's can open you online note,you must install the mongodb and mysql in your server,before this,you should see the config of the code.
```
liongo run -p 8989
```
liongo run at 8989
```
### new
```
liongo new [yourblogtitle]
```
It will generate the article file in the build file,and you can update the file to build you site
### extra
another thing is the article manager and pnote manager -> lionreact  
you can find it at:

[lionreact](https://github.com/Anteoy/lionreact)

It support the curd for articles to manage articles and nodes easy.