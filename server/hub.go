package main

// import "fmt"

// hub handles registering/unregistering of clients to channels
// and broadcast of messages to all clients in each channel.

// save copy of each document locally (for broadcast to new clients)
var contents = make(map[string]string)

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

// handles register/unregister/broadcast on channel by channel basis.
func (h *hub) run() {
  for {
    select{
    case s := <- h.register:
      // fmt.Println("wild client has appeared in the brush!")
      clients := h.channels[s.channel]
      if clients == nil {
        clients = make(map[*client]bool)
        h.channels[s.channel] = clients
      }
      h.channels[s.channel][s.client] = true
      //send the latest data for room (empty string if new room)
      s.client.send <- []byte(contents[s.channel])
    case s := <- h.unregister:
      clients := h.channels[s.channel]
      if clients != nil {
        if _, ok := clients[s.client]; ok{
          delete(clients, s.client)
          close(s.client.send)
          if len(clients) == 0 {
            delete(h.channels, s.channel)
            if len(contents[s.channel]) != 0 {
              //delete contents for channel if no more clients using it.
              delete(contents, s.channel)
            }
          }
        }
      }
    case m := <- h.broadcast:
      clients := h.channels[m.channel]
      // fmt.Println("broadcasting message to ", clients, "data is: ", m.data)
	    for c := range clients {
        select {
        case c.send <- m.data:
        contents[m.channel] = string(m.data)
        default:
          close(c.send)
          delete(clients, c)
          if len(clients) == 0 {
            delete(h.channels, m.channel)
            if len(contents[m.channel]) != 0 {
              //delete contents for channel if no more clients using it.
              delete(contents, m.channel)
            }
          }
        }
      }
    }
  }
}
