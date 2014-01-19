{{template "head.tpl"}}

{{template "top.tpl"}}
 
  		<header class="hero-unit" style="background-color:#A9F16C">
			  <table class="table table-hover table-condensed color-link">
                        <thead>
                            <tr>
                                <th>Id</th>
                                <th>a</th>
                                <th>s</th>
                              
                            </tr>
                        </thead>
                        <tbody>
                            {{range $article := .Objects}}
                            <tr>
                                <td><a href="{{$.AppUrl}}admin/article/{{$article.Id}}">{{$article.Id}}</a></td>
                    
                                <td>{{$article.Uname }}</td>
                               <td>{{$article.Tel }}</td>
                            </tr>
                            {{end}}
                        </tbody>
                    </table>

{{.Paginator.PageLinkFirst}}
<ul class="pagination pagination-sm">

    {{if .Paginator.HasPrev}}
        <li><a href="{{.Paginator.PageLinkFirst}}">首页</a></li>
        <li><a href="{{.Paginator.PageLinkPrev}}">&lt;</a></li>
    {{else}}
        <li class="disabled"><a>首页</a></li>
        <li class="disabled"><a>&lt;  </a></li>
    {{end}}
    {{range $index, $page := .Paginator.Pages}}
        <li{{if $.Paginator.IsActive .}} class="active"{{end}}>
            <a href="{{$.Paginator.PageLink $page}}">{{$page}}</a>
        </li>
    {{end}}
    {{if .Paginator.HasNext}}
        <li><a href="{{.Paginator.PageLinkNext}}">&gt;</a></li>
        <li><a href="{{.Paginator.PageLinkLast}}">末页</a></li>
    {{else}}
        <li class="disabled"><a>&gt;</a></li>
        <li class="disabled"><a>末页</a></li>
    {{end}}
</ul>
 

		</header>

{{template "foot.tpl"}}
