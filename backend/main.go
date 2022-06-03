package main

import (
	"chatting/pkg/websocket"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func routes() *mux.Router {
	pool := websocket.NewPool()
	go pool.Start()
	router := mux.NewRouter()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is up")
	})
	return router
}

func main() {
	fmt.Println("Live Chat")
	router := routes()
	http.ListenAndServe(":8080", router)
}
