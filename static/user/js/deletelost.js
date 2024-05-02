function deletelost(lostid) {
    var flag = confirm("确定删除吗，该操作不可逆")
    if (flag){
        $.ajax({
            type: "POST",
            url: "http://127.0.0.1:9526/router/DeleteLostItem",
            data: JSON.stringify({
                "lostItemId": lostid,
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
                    alert("操作成功")
                    window.location.href = 'application.html'
                }
            },
            error: function(xhr, status, error) {
                // 处理 AJAX 请求失败的情况
                new TipBox({type:'error',str:'操作失败',hasBtn:true});
                console.error("Error fetching rank data:", error);
            }
        });
    }
   
}