{{template "header"}}
<div class="top-nav">
			<ul>
			    <li><a href="/" class="on-sel">Index</li>
				<li><a href="/blog.html">Blog</a></li>
				<li><a href="/archive.html">Archive</a></li>
				<li><a href="/pages/about.html" >About</a></li>
				{{range .nav}}
				<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
				{{end}}
			</ul>
</div>
<div style="clear:both;"></div>
<div class="main">
	<div class="main-inner">
		<div id="tags-main">
			{{range $k,$v := .cats}}
        	<a href="/category.html#{{$k}}">{{$k}}</a>
       		{{end}}
		 	<div style="clear:both;"></div>
		 </div>
        <div id="tag-index">
        {{range $k,$v := .cats}}
        	<h1><a name="{{$k}}">{{$k}}</a></h1>
			{{range $v.Articles}}
            <p><a href="/articles/{{.Link}}">{{.Title}}</a></p>
            {{end}}
       	{{end}}
        </div>
		</div>
</div>
{{template "footer"}}