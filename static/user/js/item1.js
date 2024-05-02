var imgData = {}; // 定义全局变量用于存储图片数据
let admin;
$.ajax({
    "type": "POST",
    "url": "http://127.0.0.1:9526/router/GetManger",
    async: false,
    "success": function(manger) {
      admin = manger.data;
      $.ajax({
        "type": "POST",
        "url": "http://127.0.0.1:9526/router/GetLostItem",
        "success": function(data) {
            var flagfound;
            for (const a in data.data) {
                if (data.data[a].isFound == true) {
                    flagfound = "已寻得✔";
                } else {
                    flagfound = "未寻得✖";
                    if (data.data[a].foundDate != 0) {
                        flagfound = "有人上报";
                    }
                }
    
                // 发送请求获取图片数据
                $.ajax({
                    "type": "POST",
                    "url": "http://127.0.0.1:9526/router/GetLostItemImages",
                    "data": { "lostItemId": data.data[a].lostItemId },
                    "success": function(imageData) {
                   
                        // 将图片数据存储在全局变量中
                        imgData[data.data[a].lostItemId] = imageData.data;
    
                        // 渲染页面
                        renderPage(data.data[a], flagfound);
                    }
                });
            }
        }
    });
    }
});


// 渲染页面函数
function renderPage(itemData, flagfound) {
   
    console.log(admin);
    
    console.log(itemData.user);
    console.log("------");
    var truncatedUser = itemData.user.length > 22 ? itemData.user.substring(0, 30) + '...' + itemData.user.substring(40, 45): itemData.user;
    if (itemData.user == admin){
        $("#item").append(`
        <div class="col-sm-6 col-lg-4">
            <article class="product-classic" data-id="${itemData.lostItemId}">
                <div class="product-classic-header">
                    <figure class="product-classic-figure"><img src="${imgData[itemData.lostItemId]}" alt="${itemData.lostItemName}" width="270" height="200" loading="lazy"/></figure>
                    <div class="product-classic-cart"><a class="btn-cart" onclick="buy()">
                        <svg class="svg-fill" width="22" height="21">
                            <path d="M18.4141 6.76467L15.2899 0L14.0465 0.57426L16.9055 6.76467H6.03205L8.891 0.57426L7.64761 0L4.52341 6.76467H0.96875V11.6972H2.6111L4.2911 20.1385H18.6465L20.3265 11.6972H21.9688V6.76467H18.4141ZM17.5226 18.769H5.41493L4.00752 11.6971H18.93L17.5226 18.769ZM20.5992 10.3276H2.3383V8.13422H20.5992L20.5992 10.3276Z"></path>
                            <path d="M7.54916 12.7688L6.20582 13.0361L7.13263 17.693L8.47597 17.4256L7.54916 12.7688Z"></path>
                            <path d="M12.3033 12.8562H10.9337V17.6036H12.3033V12.8562Z"></path>
                            <path d="M15.6513 12.7685L14.7244 17.4254L16.0677 17.6927L16.9945 13.0359L15.6513 12.7685Z"></path>
                        </svg></a>
                    </div>
                </div>
                <div class="product-classic-body">
                    <div><input id="lostItemId" type="hidden" value="${itemData.lostItemId}"></div>
                    <h5 class="product-classic-title"><a href="single-product.html"><input id="lostItemName" type="hidden" value="${itemData.lostItemName}">#${itemData.lostItemId} ${itemData.lostItemName}</a></h5>
                    <div class="product-classic-text">${itemData.description}</div>
                    <div class="product-classic-price price">
                        <div class="price-current" ><input type="hidden" id="reward" >No Reward</div>
                    </div>
                    <div class="post-modern-meta">联系人：<a href="callme.html?address=${itemData.user}">${truncatedUser}</a></div>      
                    <ul class="btn-numbox">  
                        <li><h1 class="kucun">${flagfound}</h1></li>
                    </ul>
                    <button type="submit" class="btn" onclick="buy()" style="background-color: #8EB360"><h4>申请找回<h4></button>
                </div>
            </article>
        </div>
    `);
    };

}

   

