package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/dchest/uniuri"
  "net/http"
  "encoding/json"
)

// urls.go handles unique url assignment. 

var clientUrl string
var customChannelSet bool

// if custom url detected, sets clientUrl.
func customChannelHandler(w http.ResponseWriter, r *http.Request) {
  muxedReq := mux.Vars(r)
  clientUrl = muxedReq["channel"]
  customChannelSet = true
  http.Redirect(w, r, "/", 301)
}

// responds to client url request to use correct websocket channel.
func urlHandler(w http.ResponseWriter, r *http.Request) {
  if !customChannelSet {
    clientUrl = uniuri.NewLen(5)
  }
  customChannelSet = false
  clientUrlJSON, _ := json.Marshal(clientUrl)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  fmt.Fprintf(w, "%s", clientUrlJSON)
}
