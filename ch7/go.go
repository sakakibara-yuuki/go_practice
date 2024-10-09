//
// go.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"
import "time"

func sub() {
  for {
    fmt.Println("sub loop")
    time.Sleep(100 * time.Millisecond)
  }
}

func main() {
  go sub()
  go sub()

  for {
    fmt.Println("main loop")
    time.Sleep(200 * time.Millisecond)
  }
  
}
