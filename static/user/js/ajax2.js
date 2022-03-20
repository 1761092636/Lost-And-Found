

$.ajax({
    "type": "POST",
    "url" : "http://localhost:9526/router/products",
    "success":function(data){
        for (const a in data.data ){
            $("#best").append(`

<div class="row row-50 align-items-center">
    <div class="col-md-6">
    <img src="${data.data[a].med_img}" alt="" width="462" height="533" loading="lazy"/>
    </div>
    <div class="col-md-6">
        <div class="badge badge-lg badge-primary">区块链药品溯源</div>
        <h2>${data.data[a].med_name}</h2>
        <p class="big">${data.data[a].med_txt}</p>
        <a class="btn btn-primary" href="products.html">立即下单</a>
    </div>
</div>

               `)
        }



    }
})