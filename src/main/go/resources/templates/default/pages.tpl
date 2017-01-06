{{template "header"}}
<div class="top-nav">
	<ul>
	    <li><a href="/" class="on-sel">Index</li>
		<li><a href="/blog.html">Blog</a></li>
		<li><a href="/archive.html">Archive</a></li>
		{{range .nav}}
		<li><a href="{{.Href}}" target="{{.Target}}" >{{.Name}}</a></li>
		{{end}}
	</ul>
</div>
<div style="clear:both;"></div>
<div class="main">
	<div class="main-inner">
		 <h1>{{.p.Title}}</h1>
		 <div id="page-content">{{.p.Content|unescaped}}</div>
	</div>
</div>
{{template "footer"}}