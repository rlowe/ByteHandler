package main

import (
  "fmt"
  "tlv"
)

func main() {
  var b = []byte{ 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1 }
  fmt.Println(tlv.Handle(b))
}
