package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {

	router := gin.Default()

	router.GET("/lol", todo)
	router.GET("/ws", todo)

	// lister, _ := net.Listen("tcp", "localhost:5000")

	router.Run("localhost:8888")
}

func todo(c *gin.Context) {
	print("sdsd")

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func todoWS(c *gin.Context) {
	conn, _ := upgrader.Upgrade(c.Writer, c.Request, c.Request.Header)

	if conn != nil {
		log.Printf("dd ", conn)
	}
}
