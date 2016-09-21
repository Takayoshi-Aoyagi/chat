package main

type room struct {
	forward chan []byte
	join chan *client
	leave chan *client
	clients map[*client]bool
}

func (r * room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(coient.send)
		case msg := <-r.forward:
			// Broadcast
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send message
				default:
					// send failure
					delete(r.clients
						close(client.send)
					}
				}
			}
		}
	}
}

