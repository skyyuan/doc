<div class="footer">
    <div class="container">
        <div class="row text-center border-top">
            <span><a href="#">官方网站</a></span>
            <span>&nbsp;·&nbsp;</span>
            <span><a href="#">意见反馈</a></span>
            <span>&nbsp;·&nbsp;</span>
            <span><a href="#">项目源码</a></span>
            <span>&nbsp;·&nbsp;</span>
            <span><a href="#">使用手册</a></span>
        </div>
        {{if .site_beian}}
        <div class="row text-center">
            <a href="#">{{.site_beian}}</a>
        </div>
        {{end}}
    </div>
</div>