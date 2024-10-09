//
// if.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  a := 1
  if a == 2 {
    fmt.Println("a is 2")
  } else if a == 1 {
    fmt.Println("a is 1")
  } else {
    fmt.Println("a is not 2")
  }

  // 簡易文付きif
  if b := 100; b == 100 {
    fmt.Println("b is 100")
  }

  // if文は内部のスコープが独立している
  x := 0
  if x:=2; true {
    fmt.Println(x)
  }
  fmt.Println(x)
}
