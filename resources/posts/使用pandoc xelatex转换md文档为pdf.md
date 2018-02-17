---
date: 2018-02-16 23:33:00
title: 使用pandoc xelatex转换md文档为pdf
categories:
    - 随笔
tags:
    - pandoc markdown pdf
---
#### 环境
- 系统为ubuntu 16.04,其他linux发行版理论上可参考官方安装文档
#### 过程
1. 安装pandoc
    
    ```
    sudo apt install pandoc
    ```
2. 安装texlive-xetex（解决不能转换中文问题）
    ```
    sudo apt-get install texlive-xetex
    ```
3. 查看系统已安装的中文字体
    ```
    fc-list :lang=zh
    ```
    ```
    /usr/share/fonts/truetype/wqy/wqy-microhei.ttc: 文泉驿等宽微米黑,文泉驛等寬微米黑,WenQuanYi Micro Hei Mono:style=Regular

    ```
    这里的WenQuanYi Micro Hei Mono就是字体名
4. 把md文件转换为pdf
    ```
     pandoc -s --toc --smart --latex-engine=xelatex -V CJKmainfont='WenQuanYi Micro Hei Mono' -V geometry:margin=1in ~/桌面/resume-2018.md  -o ~/桌面/test1.pdf
    ```
   这里的-s 为standalone,--toc, --table-of-contents,CJKmainfont指定字体，我这里使用geometry:margin=1in的实际效果可以节省一页pdf.
    
### 参考
1. [https://jdhao.github.io/2017/12/10/pandoc-markdown-with-chinese/#%E5%A6%82%E4%BD%95%E5%A4%84%E7%90%86%E4%B8%AD%E6%96%87](https://jdhao.github.io/2017/12/10/pandoc-markdown-with-chinese/#%E5%A6%82%E4%BD%95%E5%A4%84%E7%90%86%E4%B8%AD%E6%96%87)
2. [https://www.zhihu.com/question/20849824](https://www.zhihu.com/question/20849824)
