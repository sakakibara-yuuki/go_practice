//
// unname_func.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  // 無名関数
  f := func(x, y int) int {
    return x + y
  }
  i := f(1, 2)
  fmt.Println(i)

  i2 := func(x, y int) int {
    return x + y
  }(1, 2)
  fmt.Println(i2)

}
