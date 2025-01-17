//
// closer.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func Later() func(string) string {
  var store string
  return func(next string) string {
    s := store
    store = next
    return s
  }
}

func main() {
  f := Later()
  fmt.Println(f("Hello"))
  fmt.Println(f("My"))
  fmt.Println(f("name"))
  fmt.Println(f("is"))
  fmt.Println(f("Golang"))
}
