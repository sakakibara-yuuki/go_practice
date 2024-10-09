//
// bool.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  fmt.Println(true && false == true)
  fmt.Println(true && true == true)
  fmt.Println(true && false == false)
  fmt.Println(true || false == true)
  fmt.Println(false || false == true)
  fmt.Println(!true)
  fmt.Println(!false)
}
