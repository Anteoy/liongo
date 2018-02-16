<html>
	<head>
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
	<meta name="description" content="{{"meta.description"|get}}"/>
	<meta name="keywords" content="{{"meta.keywords"|get}}"/>
	<meta name="author" content="{{"meta.author"|get}}"/>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"/>
	<link rel="stylesheet" href="/prettify/normalize.css"/>
	<link rel="stylesheet" href="/prettify/prettify_{{"codetheme"|get}}.css"/>
	<link rel="stylesheet" href="/css/main.css"/>
	<link rel="shortcut icon" href="/fav.ico"/>
	<script type="text/javascript" src="/prettify/prettify.js"></script>
	</head>
	<body onload="prettyPrint()" style="text-align: center">
<div class="top-nav">
	<ul>
	    <li><a href="/" class="on-sel">Index</li>
		<li><a href="/blog_1.html" >Blog</a></li>
		<li><a href="/archive.html">Date</a></li>
		<li><a href="/classify.html">Classify</a></li>
		<li><a href="/pages/about.html" >About</a></li>
		<li><a href="/pnotelogin.html" >Pnote</a></li>
		{{range .nav}}
		    <li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
		{{end}}
	</ul>
</div>
	<div style="clear:both;height:260px" id="interval"></div><!-- 中间间隔 -->
	<!-- start welcome -->
	<div style="color:white" id="welcome" class="main">
    	<p>welcome to anteoy’s site</p>

    	<p>there you can find my blog and some things about me</p>

    	<p>you can click down</p>

    	    <li><a href="/blog.html" >Blog</a></li>

    	    <li><a href="/pages/about.html" >About me</a></li>

    	<p>it’s powered by <a href="www.github.com">anteoy·liongo</a>, you can click my github to get the source code ,if you
    		find anything or some advise,please contact me,the email address is anteoy@gmail.com,you can also find me at there
    	</p>
        <p style="
               padding-top: 110px;
               margin-bottom: -40px;
           ">contact me at :</p>
    	<p>

    	    <a href="https://github.com/Anteoy"><img
                				src="/images/site/github.png"
                				alt="github" style="height:80px"/>
            </a>
            <a href="https://coding.net/u/zhoudafu"><img
                src="/resources/images/site/coding.png"
                alt="github" style="height:80px"/>
            </a>
            <a href="http://blog.csdn.net/yan_chou"><img
                    src="/resources/images/site/csdn.png"
                    alt="github" style="height:80px"/>
            </a>
            <a href="https://twitter.com/AnteoyChou"><img
                src="/resources/images/site/twitter.png"
                alt="github" style="height:90px"/>
            </a>
    		<!-- github<a href="https://github.com/Anteoy"><img
    				src="https://raw.githubusercontent.com/Anteoy/liongo/dev/src/main/go/resources/pictures/github.png"
    				alt="github" style="height:80px"/>
    		</a>
    		<a href="https://coding.net/u/zhoudafu"><img
                src="https://raw.githubusercontent.com/Anteoy/liongo/dev/src/main/go/resources/pictures/coding1.png"
                alt="github" style="height:80px"/>
            </a>
            <a href="http://blog.csdn.net/yan_chou"><img
                    src="https://raw.githubusercontent.com/Anteoy/liongo/dev/src/main/go/resources/pictures/csdn.png"
                    alt="github" style="height:80px"/>
            </a>
            <a href="https://twitter.com/AnteoyChou"><img
                src="https://raw.githubusercontent.com/Anteoy/liongo/dev/src/main/go/resources/pictures/twitter.jpg"
                alt="github" style="height:80px"/>
            </a> -->
    	</p>
    </div>

{{template "footer"}}