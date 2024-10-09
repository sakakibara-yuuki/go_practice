//
// slice.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  // 配列と異なり, 要素数を指定しない
  var sl []int
  fmt.Println(sl)

  //明示的な宣言
  var sl2 = []int{100, 200}
  fmt.Println(sl2)

  //暗黙的な宣言
  sl3 := []int{100, 200, 300}
  fmt.Println(sl3)

  sl4 := make([]int, 5)
  fmt.Println(sl4)

  sl2[0] = 1000
  fmt.Println(sl2)

  sl5 := []int{1, 2, 3, 4, 5}
  fmt.Println(sl5[0])
  fmt.Println(sl5[2:4])
  fmt.Println(sl5[:2])
  fmt.Println(sl5[2:])
  fmt.Println(sl5[:])
  fmt.Println(sl5[len(sl5)-1])
  fmt.Println(sl5[1:len(sl5)-1])
}
