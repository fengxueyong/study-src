package main

import "fmt"
import "os"

func main() {
  fmt.Println(os.Args)
  fmt.Println("hello world")
  // return 1
  os.Exit(0)
}
