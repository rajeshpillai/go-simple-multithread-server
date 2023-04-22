package main

import (
  "fmt"
  "net"
  "log"
)

func main () {
  listener, err := net.Listen("tcp", ":1729")
  if err != nil {
    log.Fatal(err)
  }

  conn, err :=  listener.Accept()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(conn)
}
