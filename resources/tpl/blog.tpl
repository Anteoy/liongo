{{template "header"}}
<div class="top-nav">
	<ul>
	    <li><a href="/" >Index</li>
		<li><a href="/blog_1.html" class="on-sel">Blog</a></li>
		<li><a href="/archive.html">Date</a></li>
		<li><a href="/classify.html">Classify</a></li>
		<li><a href="/pages/about.html" >About</a></li>
		<li><a href="/pnotelogin.html" >Pnote</a></li>
		{{range .nav}}
		<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
		{{end}}
	</ul>
</div>

<div style="clear:both;height:50px" id="interval"></div><!-- 中间间隔 -->
<div class="main">
	<div class="main-inner">
		<div class="article-list">
            {{range .ar}}
                <div class="article">
                    <p class="title"><a href="/articles/{{.Link}}">{{.Title}}</a></p>
                    <p class="abstract">&lt;摘要&gt;: {{.Abstract}}&nbsp;&nbsp;<a href="/articles/{{.Link}}">Read more</a></p>
                    <p class="meta">Author {{.Author}} | Posted {{.Date}} | Tags
                    {{range .Tags}}
                    <a class="tag" href="/tag.html#{{.Name}}">{{.Name}}</a>
                    {{end}}
                    </p>
                </div>
            {{end}}
		</div>
		<div style="margin-bottom:200px;">
            <ul class="pager main-pager">
                <li class="previous" style="display:{{.display0}}">
                    <a href="{{.pre}}">← 上一页</a>
                </li>
                <li class="next" style="display:{{.display1}}">
                    <a href="{{.next}}">下一页 →</a>
                </li>
                <li class="next" >
                    <input id="page_id" style="width: 21px;"></input><a href="javascript:void(0);" onclick="jump()">输入页码，点击跳转</a>
                </li>
                <li class="next">
                    <div>当前第{{.i}}页;共{{.total}}页</div>
                </li>
            </ul>
        </div>
	</div>
</div>
<script type="text/javascript">
    function jump(){
        var page_id=document.getElementById('page_id').value
        if (page_id==0 || page_id == undefined){
            alert("请输入需要跳转的页码，然后再次点击跳转")
        }
        console.log("page_id:", page_id)
        window.location = "{{.jump_url}}" + page_id + ".html";
    }
</script>
{{template "footer"}}