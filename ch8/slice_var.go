//
// slice_var.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

// カンマを入れないように注意
// あるしゅ...intという型
func Sum(s ...int) int {
  n := 0
  for _, v := range s {
    n += v
  }
  return n
}

func main() {
  fmt.Println(Sum(1, 2, 3))
  fmt.Println(Sum(1, 2, 3, 4, 5))
  fmt.Println(Sum())

  sl := []int{1, 2, 3}
  // スライスを渡す場合は...をつける
  fmt.Println(Sum(sl...))
}
