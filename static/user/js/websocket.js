var ws = new WebSocket('ws://127.0.0.1:9526/router/ws');  
var messageContainer = document.getElementById('message-container');  
var isRealTimeMessage = false; // 用来跟踪是否收到实时消息  
  
ws.onopen = function(event) {  
    console.log("WebSocket connection is open now.");  
    // 这里可以发送一些初始化消息给服务器，但不是实时信息  
};  
  
ws.onmessage = function(event) {  
    var message = event.data;  
    console.log("Received message: " + message);  
      
    // 检查消息是否是实时信息（这里假设实时信息以"REALTIME:"开头）  
    if (message.startsWith("REALTIME:")) {  
        isRealTimeMessage = true;  
        alert("Received real-time message: " + message);  
    } else {  
        // 对于非实时信息，只添加到消息容器中，不弹窗  
        var messageElement = document.createElement('div');  
        messageElement.textContent = message;  
        messageContainer.appendChild(messageElement);  
        messageContainer.scrollTop = messageContainer.scrollHeight;  
    }  
      
    // 重置标识，确保下一次收到消息时不会错误地弹窗  
    isRealTimeMessage = false;  
};  
  
ws.onerror = function(error) {  
    console.error("WebSocket error observed:", error);  
};  
  
// 如果您想要处理连接关闭事件，可以取消注释下面的代码  
// ws.onclose = function(event) {  
//     console.log("WebSocket connection closed.");  
// };