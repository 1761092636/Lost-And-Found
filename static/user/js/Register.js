
    $("#Register_bt").click(function(a){
a = checkPW()
        console.log(a)
        if (a !== false){
            $.ajax({
                "type" : "POST",
                "url" : "http://127.0.0.1:9526/router/Register",
                xhrFields: {  
                    withCredentials: true // 允许携带凭证  
                },  
                crossDomain: true, // 允许跨域请求 
                data :{
                    "user_name":$("#contact-name").val(),
                    "name":$("#name").val(),
                    "phone":$("#userPhone").val(),
                 
                    "user_password":$("#contact-email").val()},
                success:function(data){
                    if (data.code === 0){
                        // confirm("注册成功，去登陆?")

                        new TipBox({type:'success',str:'注册成功',hasBtn:true});
                        window.location.href = "login.html"

                    }
                    console.log(data.data)
                }

            })
        }


    })


