//
// map.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  var m = map[string]int{"A": 100, "B": 200}
  fmt.Println(m)

  // 暗黙的型宣言
  m2 := map[string]int{"A": 100, "B": 200}
  fmt.Println(m2)

  // 改行して宣言
  m3 := map[int]string{
    1: "A",
    2: "B",
  }
  fmt.Println(m3)

  //make関数を使って宣言
  m4 := make(map[int]string)
  fmt.Println(m4)

  m4[1] = "Japan"
  m4[2] = "USA"
  fmt.Println(m4)

  fmt.Println(m["A"])
  fmt.Println(m4[2])
  fmt.Println(m4[3]) //本当は存在しないが、初期値が返る

  // 取り出しに成功したか判断する
  s, ok := m4[3]
  if !ok {
    fmt.Println("error")
  }
  fmt.Println(s, ok)

  m4[2] = "US"
  fmt.Println(m4)

  m4[3] = "China"
  fmt.Println(m4)

  // 削除
  delete(m4, 3)
  fmt.Println(m4)

  // len関数
  fmt.Println(len(m4))
}
