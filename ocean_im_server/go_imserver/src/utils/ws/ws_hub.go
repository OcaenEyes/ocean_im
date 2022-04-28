package ws

type MyWsHub struct {
	Clients      map[*MyWsClient]bool
	Broadcast    chan []byte
	ChRegister   chan *MyWsClient
	ChUnregister chan *MyWsClient
}

func NewMyWsHub() *MyWsHub {
	return &MyWsHub{
		Broadcast:    make(chan []byte),
		ChRegister:   make(chan *MyWsClient),
		ChUnregister: make(chan *MyWsClient),
		Clients:      make(map[*MyWsClient]bool),
	}
}

func (m *MyWsHub) Run() {
	for {
		select {
		case client := <-m.ChRegister:
			m.Clients[client] = true
		case client := <-m.ChUnregister:
			if _, ok := m.Clients[client]; ok {
				delete(m.Clients, client)
				close(client.Send)
			}
		case message := <-m.Broadcast:
			for client := range m.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(m.Clients, client)
				}
			}
		}
	}
}
