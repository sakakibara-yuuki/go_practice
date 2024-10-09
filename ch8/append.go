//
// append.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  sl := []int{100, 200}
  fmt.Println(sl)

  // append
  sl = append(sl, 300)
  fmt.Println(sl)

  // 可変長
  sl = append(sl, 400, 500)
  fmt.Println(sl)

  // make 要素数を指定してスライスを作成
  sl2 := make([]int, 5)
  fmt.Println(sl2)
  fmt.Println(len(sl2))
  fmt.Println(cap(sl2))

  sl3 := make([]int, 5, 10)
  fmt.Println(sl3)
  fmt.Println(len(sl3))
  fmt.Println(cap(sl3))

  sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
  fmt.Println(len(sl3))
  fmt.Println(cap(sl3))

}
