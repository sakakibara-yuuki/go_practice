//
// init.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func init() {
  fmt.Println("init")
}

func init() {
  fmt.Println("init2")
}

func main() {
  fmt.Println("main")
}
