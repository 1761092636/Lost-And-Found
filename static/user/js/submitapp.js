
  function buy(index){
    console.log(index)
    a = confirm("提交拾得申请？")
    if (a === false){
      console.log("用户取消操作")

    }else {


      $.ajax({
        type: "POST",
        url: "http://127.0.0.1:9526/router/SubmitApplication",
        async: false,
         xhrFields: {  
        withCredentials: true // 允许携带凭证  
    },  
        crossDomain: true, // 允许跨域请求  
        data: {
          "lost_item_id": index,
        },
        success: function (data) {

          if (data.code === 0){

            new TipBox({type:'success',str:'申请成功',hasBtn:true});

            $(".okoButton").click(function (){
              window.location.href = 'index.html'
            })
            return data.data
          }else {

            new TipBox({type:'error',str:'申请失败，请检查余额是否充足',hasBtn:true});
            $(".okoButton").click(function (){

            })
          }
          console.log(data)

        }
      })



    }

  }
