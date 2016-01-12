{{define "navbar"}}
        <meta http-equiv="CONTENT-TYPE" content="text/html"; charset="utf-8">
		<div class="navbar navbar-default navbar-fixed-top">
			<div class="container">
				<a class="navbar-brand" href="/">My Blog</a>
				<ul class="nav navbar-nav">
					<li {{if .IsHome}}class="active"{{end}}><a href="/">Header</a> </li>
					<li {{if .IsCategory}}class="active"{{end}}><a href="/category">Category</a> </li>
					<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">Topics</a> </li>
				</ul>
			</div>
		</div>
{{end}}