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

  r := mux.NewRouter()

  r.HandleFunc("/ws/{channel}", serveWS)
  // r.HandleFunc("/db", saveDoc).Methods("POST")
  // r.HandleFunc("/db", retrieveDoc).Methods("GET")
  // r.HandleFunc("/db", updateDoc).Methods("UPDATE")

  // Serve static files
  r.PathPrefix("/").Handler(http.FileServer(http.Dir("../client")))
  // handle unique urls
  //r.HandleFunc("/" {id:param}, urlhandler)

  log.Fatal(http.ListenAndServe(":8000", r))
}
