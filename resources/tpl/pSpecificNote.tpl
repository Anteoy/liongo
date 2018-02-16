<!DOCTYPE html>
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
	<body onload="prettyPrint()">

        <div class="top-nav">
        			<ul>
        			    <li><a href="/" >Index</li>
        				<li><a href="/blog_1.html" >Blog</a></li>
        				<li><a href="/archive.html">Date</a></li>
        				<li><a href="/classify.html">Classify</a></li>
        				<li><a href="/pages/about.html">About</a></li>
        				<li><a href="/pnotelogin.html" >Pnote</a></li>
        				{{range .Nav}}
        				<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
        				{{end}}
        			</ul>
        </div>
        <div style="clear:both;"></div>
        <div class="main">
        	<div class="main-inner">
        		 <div id="article-content"> {{.fi.Content|unescaped}} </div>
        		<hr/>
        	</div>
        </div>
        </div>
        <script type="text/javascript" src="/js/jquery.js"></script>

        <div id="footer">
            <div id="footer-inner" style="bottom: 0;position: fixed;right: 0;left: 0;">
                <p id="copyright">Copyright (c) {{"copyright.beginYear" | get}} - {{"copyright.endYear" | get}} {{"copyright.owner"|get}} &nbsp;
                Powered by <a href="https://github.com/Anteoy/liongo">liongo</a>
                </p>
            </div>
        </div>
    </body>
</html>
