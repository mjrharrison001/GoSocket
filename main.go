package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
)

// switch protocol (http -> websockets)
var upgrader = websocket.Upgrader{
  ReadBufferSize:   1024,
  WriteBufferSize:  1024,
  CheckOrigin:      func(r *http.Request) bool { return true }, //allow requests from any origin (for now)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  //fmt.Fprintf(w, "Hello from go")
  var socket *websocket.Conn
  var err error
  socket, err = upgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Println(err)
    return
  }
  for {
    // := declared a var and assigned the type returned
    msgType, msg, err := socket.ReadMessage()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(string(msg))
    if err = socket.WriteMessage(msgType, msg); err != nil {
      fmt.Println(err)
      return
    }

  }
}
