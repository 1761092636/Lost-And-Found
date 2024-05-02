  //验证手机号 
  function checkPhone(){ 
    var userphone = document.getElementById('userPhone'); 
    var phonrErr = document.getElementById('phoneErr'); 
    var pattern = /^1[34578]\d{9}$/; //验证手机号正则表达式 
    if(!pattern.test(userphone.value)){
      phonrErr.innerHTML="手机号码不合规范" 
      phonrErr.className="error" 
      return false; 
      } 
     else{ 
       phonrErr.innerHTML="OK"
       phonrErr.className="success"; 
       return true; 
       } 
    } 