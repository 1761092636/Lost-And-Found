function confirmapp(index, lost_item_id, application_id) {
    $.ajax({
        type: "POST",
        url: "http://127.0.0.1:9526/router/ConfirmApplication",
        data: JSON.stringify({
            "lost_item_id": lost_item_id,
            "application_id": application_id,
            "result": index
        }),
        dataType: "json",
        contentType: "application/json", // 设置 content-type
        xhrFields: {
            withCredentials: true // 允许携带凭证  
        },
        crossDomain: true, // 允许跨域请求
        success: function(data) {
            // 处理成功的回调函数
            if (data.code === 0){
                alert("ok")
                console.log("溯源码:",data.data)
                $("#personal_order").append(
                    `
                    <div class="post-modern-meta" style="margin-bottom: 10px;">溯源码：<a href="#">${data.data}</a></div>
                    `
                );
                new TipBox({type:'success',str:'操作成功',hasBtn:true});
            }
        },
        error: function(xhr, status, error) {
            // 处理 AJAX 请求失败的情况
            new TipBox({type:'error',str:'操作失败',hasBtn:true});
            console.error("Error fetching rank data:", error);
        }
    });
}
