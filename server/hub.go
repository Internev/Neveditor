package main

type message struct {
  data []byte
  channel string
}

type subscription struct {
  client *client
  channel string
}

type hub struct {
  channels map[string]map[*client]bool
  broadcast chan message
  register chan subscription
  unregister chan subscription
}

var h = hub{
  broadcast:  make(chan message),
  register:   make(chan subscription),
  unregister: make(chan subscription),
  channels:   make(map[string]map[*client]bool),
}

func (h *hub) run() {
  for {
    select{
    case s := <- h.register:
      clients := h.channels[s.channel]
      if clients == nil {
        clients = make(map[*client]bool)
        h.channels[s.channel] = clients
      }
      h.channels[s.channel] = true
    case s := <- h.unregister:
      clients := h.channels[s.channel]
      if clients != nil {
        if _, ok := clients[s.client]; ok{
          delete(clients, s.client)
          close(s.client.send)
          if len(clients) == 0 {
            delete(h.channels, s.channel)
          }
        }
      }
    case m := <- h.broadcast:
      clients := h.channels[m.channel]
	    for c := range clients {
        select {
        case c.send <- m.data:
        default:
          close(c.send)
          delete(clients, c)
          if len(clients) == 0 {
            delete(h.channels, m.channel)
          }
        }
      }
    }
  }
}
