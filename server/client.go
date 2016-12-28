package main

import (
  "log"
  "net/http"
  "time"
  "github.com/gorilla/websocket"
  "github.com/gorilla/mux"
)

const (
  writeWait = 10 * time.Second
  pongWait = 60 * time.Second
  pingPeriod = (pongWait * 9) / 10
  // This is for max size, check here if we have problems on large Docs!
  maxMessageSize = 1024 * 1024
)

// sets up Upgrade object with methods for taking a normal HTTP connection and upgrading it to a websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize: maxMessageSize,
	WriteBufferSize: maxMessageSize,
}

// sets up a struct for client connections

type client struct {
	//ws connection
	ws *websocket.Conn
	//buffered channel that receives data type bytes
	send chan []byte
}

// sends message struct (msg, room) to the hub for broadcast
func (s subscription) readPump() {
    c := s.client
    defer func() {
        h.unregister <- s
        c.ws.Close()
    }()
    c.ws.SetReadLimit(maxMessageSize)
    c.ws.SetReadDeadline(time.Now().Add(pongWait))
    c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
    for {
        _, msg, err := c.ws.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
                log.Printf("error: %v", err)
            }
            break
        }
        m := message{msg, s.channel}
        h.broadcast <- m
    }
}

// Sets deadline for write attempt, at 10secs gives up and throws error.
func (c *client) write(mt int, payload []byte) error {
  c.ws.SetWriteDeadline(time.Now().Add(writeWait))
  return c.ws.WriteMessage(mt, payload)
}

func (s *subscription) writePump() {
	c := s.client
	ticker := time.NewTicker(pingPeriod)
	defer func(){
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err:= c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <- ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	vars := mux.Vars(r)
	log.Println("this is the channel: ", vars["channel"])
	if err != nil {
		log.Println("you have failed in serving the WS master", err)
		return
	}

	c := &client{send: make(chan []byte, maxMessageSize), ws:ws}
	s := subscription{c, vars["channel"]}
	h.register <- s
	go s.writePump()
	//why is readPump not also a go subroutine?
	s.readPump()
}
