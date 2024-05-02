  //验证手机号 
  function checkPW(){
    var userphone = document.getElementById('contact-email');
    var phonrErr = document.getElementById('pwErr');
    var pattern = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,20}$/; //验证手机号正则表达式
    if(!pattern.test(userphone.value)){ 
      phonrErr.innerHTML="密码不合规范"
      phonrErr.className="error"
      return false; 
      } 
     else{ 
       phonrErr.innerHTML="OK" 
       phonrErr.className="success"; 
       return true; 
       } 
    }
