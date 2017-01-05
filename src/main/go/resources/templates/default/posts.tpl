{{template "header"}}
<div class="top-nav">
			<ul>
				<li><a href="/" >Index</a></li>
				<li><a href="/tag.html">Tags</a></li>
				<li><a href="/category.html">Categories</a></li>
				<li><a href="/archive.html">Archive</a></li>
				{{range .Nav}}
				<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
				{{end}}
				<li><a href="/rss.xml" class="rss" title="feed"></a></li>
			</ul>
</div>
<div style="clear:both;"></div>
<div class="main">
	<div class="main-inner">
		 <div id="article-title">
		 	<a href="/{{.fi.Link}}">{{.fi.Title}}</a>
		 </div>
		 <div id="article-meta">Author {{.fi.Author}}  | Posted {{.fi.Date}} </div>

		  <div id="article-tags">
		  {{range .Tags}}
		  <a class="tag" href="/tag.html#{{.Name}}">
		  {{.Name}}</a> 
		  {{end}}
		  </div>
		 <div id="article-content"> {{.fi.Content|unescaped}} </div>
		<hr/>
	</div>
</div>

</div>

<script type="text/javascript" src="/assets/themes/{{"theme"|get}}/jquery.js"></script>

{{template "footer"}}