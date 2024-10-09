//
// panic.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  defer func() {
    if x := recover(); x != nil {
      fmt.Println("Recover from panic")
    }
  }()
  panic("runtime error")
  fmt.Println("Start")
}
