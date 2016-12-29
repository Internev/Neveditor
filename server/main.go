package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/gorilla/mux"
)

func main() {
  fmt.Println("Go server, go! (8k)")

  // run the hub to start websockets
  go h.run()

  r := mux.NewRouter()

  r.HandleFunc("/ws/{channel}", serveWS)
  // r.HandleFunc("/db", saveDoc).Methods("POST")
  // r.HandleFunc("/db", retrieveDoc).Methods("GET")
  // r.HandleFunc("/db", updateDoc).Methods("UPDATE")
  // This is a terrible regex, mux does not support lots of regex stuff :(
  // Matches 5 char unique url.
  r.HandleFunc("/{channel:[0-9A-Za-z][0-9A-Za-z][0-9A-Za-z][0-9A-Za-z][0-9A-Za-z]}", customChannelHandler)

  r.HandleFunc("/getUrl", urlHandler).Methods("GET")

  // Serve static files
  r.PathPrefix("/").Handler(http.FileServer(http.Dir("../client")))
  // handle unique urls
  //r.HandleFunc("/" {id:param}, urlhandler)

  log.Fatal(http.ListenAndServe(":8000", r))
}
