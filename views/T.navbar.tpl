{{define "navbar"}}
<a class="navbar-brand" href="/">My Blog-我的博客</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if .IsHome}}class="active"{{end}}><a href="/">Header-首页</a> </li>
		<li {{if .IsCategory}}class="active"{{end}}><a href="/category">Category-分类</a> </li>
		<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">Topics-文章</a> </li>
	</ul>
</div>			

<div class="pull-right">
	<ul class="nav navbar-nav">
		{{if .IsLogin}}
			<li><a href="/login?exit=true">Exit-退出</a></li>
		{{else}}
			<li><a href="/login">Login-登陆</a></li>
		{{end}}
	</ul>
</div>
{{end}}