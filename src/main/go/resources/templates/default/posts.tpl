{{template "header"}}
<div class="top-nav">
			<ul>
			    <li><a href="/" >Index</li>
				<li><a href="/blog.html" >Blog</a></li>
				<li><a href="/archive.html">Date</a></li>
				<li><a href="/classify.html">Classify</a></li>
				<li><a href="/pages/about.html">About</a></li>
				{{range .Nav}}
				<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
				{{end}}
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