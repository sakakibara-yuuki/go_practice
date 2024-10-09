//
// for.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  // i := 0
  // for {
  //   i++
  //   if i == 3 {
  //     break
  //   }
  //   fmt.Println("loop")
  // }

  // point := 0
  // for point < 10 {
  //   fmt.Println(point)
  //   point++
  // }

  // for i:=0; i < 10; i++ {
  //   if i == 3 {
  //     continue
  //   }
  //   if i == 6 {
  //     break
  //   }
  //   fmt.Println(i)
  // }

  // 配列の要素を取り出す
  // arr := [3]int{1, 2, 3}
  // for i:=0; i < len(arr); i++ {
  //   fmt.Println(arr[i])
  // }

  // 範囲式のfor
  // arr := [3]int{1, 2, 3}
  // for i, v:= range arr {
  //   fmt.Println(i, v)
  // }

  // sl := []string{"Python", "PHP", "Golang"}
  // for i, v:= range sl {
  //   fmt.Println(i, v)
  // }

  // map
  m := map[string]int{"apple": 100, "banana": 200}
  for k, v:= range m {
    fmt.Println(k, v)
  }

}
