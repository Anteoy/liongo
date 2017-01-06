{{template "header"}}
<div class="top-nav">
			<ul>
			    <li><a href="/" class="on-sel">Index</li>
                <li><a href="/blog.html" >Blog</a></li>
                <li><a href="/archive.html" class="on-sel">Archive</a></li>
                {{range .nav}}
                <li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
                {{end}}
                <!-- <li><a href="/rss.xml" class="rss" title="feed"></a></li> -->
            </ul>
</div>
<div style="clear:both;" id="interval"></div>
<div class="main">
	<div class="main-inner">
        <div id="tag-index">
        {{range .archives}}
        	<h1>{{.Year}}</h1>
			{{range .Months}}
				<h2>{{.Month}}</h2>
				{{range .Articles}}
           			<p><a href="/articles/{{.Link}}">{{.Title}}</a></p>
           		{{end}}
            {{end}}
       	{{end}}
        </div>
		</div>
</div>
{{template "footer"}}