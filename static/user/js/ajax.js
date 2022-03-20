//新闻资讯
$.ajax({
    "type": "POST",
    "url" : "http://localhost:9526/router/news",
    "success":function(data){
        for (const a in data.data ){
            $("#news").append(`
 <article class="post-modern" >
               <figure class="post-modern-figure"><a class="link-image" href="single-post.html"><img  src="${data.data[a].news_img}" alt="" width="760" height="450" loading="lazy"/></a>
                
                </figure>
                <div class="post-modern-body">
                  <div class="row row-20 justify-content-between">
                    <div class="col-auto">
                      <div class="post-modern-meta">作者： <a href="#">${data.data[a].author}</a>  / ${data.data[a].timeset} </div>
                    </div>

                  </div>
                  <h3 class="post-modern-title"><a href="single-post.html">${data.data[a].title}</a></h3>
                  <p class="big post-modern-text">${data.data[a].content}</p>
                </div>
                <div class="post-modern-footer">
                  <div class="row row-20 justify-content-between align-items-center">
                    <div class="col-auto"><a class="btn btn-secondary" href="single-post.html">阅读更多</a></div>

                  </div>
                </div>
               </article>
               
               `)
        }



    }
})