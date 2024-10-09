//
// alge.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  fmt.Println(1 + 3)
  fmt.Println(1 - 3)
  fmt.Println(5 * 4)
  fmt.Println(5 / 4)
  fmt.Println(5 % 4)
  fmt.Println("ABC" + "DEF")
  n := 0
  n += 4
  fmt.Println(n)
  n++
  fmt.Println(n)
  n--
  fmt.Println(n)

  s := "ABC"
  s += "DEF"
  fmt.Println(s)
}
