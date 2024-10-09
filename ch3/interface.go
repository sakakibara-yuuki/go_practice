//
// interface.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func main() {
  // あらゆる型と互換性のあるインターフェース型
  var x interface{}
  fmt.Println(x)

  x = 1
  fmt.Println(x)
  x = "hee"
  fmt.Println(x)
  x = 3.1
  fmt.Println(x)
  x = []int{1, 2, 3}
  fmt.Println(x)

  // データ特有の演算ができない
  x = 2
  fmt.Println(x + 3)
}
