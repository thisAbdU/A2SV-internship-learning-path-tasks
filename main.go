package main

import (
	"fmt"
	"strings"
)

func main() {
  word := "hello"
  var list = strings.Split(word, "")
  fmt.Println(list)
}