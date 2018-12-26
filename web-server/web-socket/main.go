package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space = []byte{' '}
)

func Echo(ws *websocket.Conn) {
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//var reply string
		//if err = websocket.Message.Receive(ws, &reply); err != nil {
		//	fmt.Println("Can't receive")
		//	break
		//}
		//fmt.Println("Received back from client: " + reply)
		//
		//msg := "Received: " + reply
		//fmt.Println("Sending to client: " + msg)
		//
		//if err = websocket.Message.Send(ws, msg); err != nil {
		//	fmt.Println("Can't Send")
		//	break
		//}
	}
}
func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
