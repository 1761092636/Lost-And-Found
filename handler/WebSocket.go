package handler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 用户与WebSocket连接的映射
	clientConnections = make(map[string]*websocket.Conn)
	mu                sync.Mutex // 用于保护clientConnections的互斥锁
)

// 处理WebSocket连接
func serveWebsocket(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("username").(string) // 假设session中存储了user_id
	if userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing user_id in session"})
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket:", err)
		return
	}

	mu.Lock()
	clientConnections[userID] = ws
	mu.Unlock()

}

// SendMessageToUser 向指定用户发送WebSocket消息
func SendMessageToUser(userID string, message string) error {
	mu.Lock()
	defer mu.Unlock()

	conn, exists := clientConnections[userID]
	if !exists {
		fmt.Printf("User %s not found for sending message\n", userID)
		return nil
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Printf("Error sending message to user %s: %v\n", userID, err)
		delete(clientConnections, userID) // 从映射中删除连接，因为可能已经断开
	}
	return nil
}

// 假设这是在其他地方调用，通知特定用户的函数
func notifyUser(username string, message string) {

}
