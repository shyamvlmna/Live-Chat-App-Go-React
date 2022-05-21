package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    map[*Client]bool{},
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}

		case Message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in a Pool")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(Message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
