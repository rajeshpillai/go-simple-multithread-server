package main

import (
  "fmt"
  "net"
)

func main () {
  net.Listen("tcp", ":1729")
  fmt.Println("Helloo, World!")
}
