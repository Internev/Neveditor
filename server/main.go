package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/gorilla/mux"
)

// handles routing
func main() {
  fmt.Println("Go server, go! (8k)")

  // run the hub to start websockets
  go h.run()

  r := mux.NewRouter()

  r.HandleFunc("/ws/{channel}", serveWS)

  // db methods, for when db implemented.
  // r.HandleFunc("/db", saveDoc).Methods("POST")
  // r.HandleFunc("/db", retrieveDoc).Methods("GET")
  // r.HandleFunc("/db", updateDoc).Methods("UPDATE")

  // This is a terrible regex, mux does not support lots of regex stuff :(
  // Matches 5 char unique url if provided.
  r.HandleFunc("/{channel:[0-9A-Za-z][0-9A-Za-z][0-9A-Za-z][0-9A-Za-z][0-9A-Za-z]}", customChannelHandler)

  // special route for new connection, must ask /getUrl for unique url identifier for channel.
  r.HandleFunc("/getUrl", urlHandler).Methods("GET")

  // Serve static files - nb: this is dependent on run location (ie: it's set up to be run from root)
  r.PathPrefix("/").Handler(http.FileServer(http.Dir("client")))

  log.Fatal(http.ListenAndServe(":8000", r))
}
