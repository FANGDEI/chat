package web

var (
	MyHub *Hub
)

func init() {
	MyHub = NewHub()
	go MyHub.Run()
}

type Hub struct {
	//mutex      *sync.Mutex
	Clients map[int64]*Client
	//Broadcast  chan []byte
	Register   chan *Client
	UnRegister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		//mutex:      &sync.Mutex{},
		Clients: make(map[int64]*Client),
		//Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	defaultLogger.Info("hub run")
	for {
		select {
		case conn := <-h.Register:
			h.Clients[conn.UserID] = conn
		case conn := <-h.UnRegister:
			if _, ok := h.Clients[conn.UserID]; ok {
				delete(h.Clients, conn.UserID)
			}
		}
	}
}
