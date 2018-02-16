{{template "header"}}
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
<div style="clear:both;height:50px" id="interval"></div><!-- 中间间隔 -->
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
		<div id="SOHUCS" ></div>
        <script charset="utf-8" type="text/javascript" src="https://changyan.sohu.com/upload/changyan.js" ></script>
        <script type="text/javascript">
            window.changyan.api.config({
            appid: 'cytroRXMV',
            conf: 'prod_f03f628ad1f20646fa84b64386060ea1'
            });
        </script>
	</div>
</div>

</div>

<script type="text/javascript" src="/js/jquery.js"></script>

{{template "footer"}}