$("#login_bt").click(function(a){
    a = checkPW()
    console.log(a)
        $.ajax({
            "type":"POST",
            "url":"http://127.0.0.1:9526/router/Login",
            xhrFields: {  
                withCredentials: true // 允许携带凭证  
            },  
            crossDomain: true, // 允许跨域请求  
            data:{"user_name":$("#contact-name").val(),
                "user_password":$("#contact-email").val()},

            success:function(data){
                if (data.code === 0){

                    new TipBox({type:'success',str:'登录成功',hasBtn:true});

                    $(".okoButton").click(function (){
                   window.location.href = 'index.html'
      })
                    return data.data
                }else {

                    new TipBox({type:'error',str:'请检查用户名或密码',hasBtn:true});
                    $(".okoButton").click(function (){

                    })
                }
                console.log(data);



            }

        })

    }

)