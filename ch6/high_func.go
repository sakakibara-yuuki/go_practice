//
// high_func.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func ReturnFunc() func() {
  return func() {
    fmt.Println("I'm a function")
  }
}

func CallFunction(f func()) {
  f()
}

func main() {
  f := ReturnFunc()
  f()

  CallFunction(func() {
    fmt.Println("I'm a function")
  })
}
