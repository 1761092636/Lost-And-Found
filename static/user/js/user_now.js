$("#user_now").click(function (data){
    $.ajax({
      "type": "GET",
      "url" : "http://127.0.0.1:9526/router/GetUser",
      xhrFields: {  
          withCredentials: true // 允许携带凭证  
      },  
      crossDomain: true, // 允许跨域请求  
      "success":function(data) {
        if (data === ""){
          window.location.href = "login.html"
  
        }
        else {
           window.location.href = "person.html"
        }
  
  
      }
  
  
    })
  });