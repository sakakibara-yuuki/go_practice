//
// copy.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  // 参照型
  // 参照渡しになっている
  // sl := []int{100, 200}
  // sl2 := sl
  // sl2[0] = 1000
  // fmt.Println(sl)
  //
  // var i int = 10
  // i2 := i
  // i2 = 100
  // fmt.Println(i, i2)

  sl := []int{1, 2, 3, 4, 5, 6}
  sl2 := make([]int, 5, 10)
  fmt.Println(sl2)
  n := copy(sl2, sl)
  fmt.Println(n, sl2)
}
