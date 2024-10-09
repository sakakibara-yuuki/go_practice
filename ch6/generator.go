//
// generator.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func integers() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
  ints := integers()
  fmt.Println(ints())
  fmt.Println(ints())
  fmt.Println(ints())

  otherints := integers()
  fmt.Println(otherints())
  fmt.Println(otherints())
  fmt.Println(otherints())
  fmt.Println(otherints())
}
