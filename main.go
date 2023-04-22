package main

import (
  "net"
  "log"
  "time"
)

func do (conn net.Conn) {
  
  buf  := make([]byte, 1024)
  _, err :=  conn.Read(buf)  // blocking call
  if err != nil {
    log.Fatal(err)
  }

  //Processing
  log.Println("Processing the request...")
  // 8 seconds
  time.Sleep(8 * time.Second)  // Long running simulation

  conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, Worrld!\r\n"))
  conn.Close()
}

func main () {
  listener, err := net.Listen("tcp", ":1729")
  if err != nil {
    log.Fatal(err)
  }

  for {
    log.Println("Waiting for a client to connect...")
    conn, err :=  listener.Accept()
    if err != nil {
      log.Fatal(err)
    }

    log.Println("Client connected!")
    go do (conn) // go routine
  }
}
