//热门新闻
$.ajax({
    "type": "POST",
    "url" : "http://localhost:9526/router/news",
    "success":function(data){
        for (const a in data.data ){
            $("#hot").append(`

                 <div class="post-minimal">
                        <figure class="post-minimal-figure"><a class="link-image" href="#"><img src="${data.data[a].news_img}" alt="" width="80" height="80" loading="lazy"/></a></figure>
                        <div class="post-minimal-body">
                          <div class="post-minimal-title"><a href="#">${data.data[a].title}</a></div>
                          <div class="post-minimal-time">
                            <time datetime="2021">${data.data[a].timeset}</time>
                          </div>
                        </div>
                      </div>
               `)
        }



    }
})