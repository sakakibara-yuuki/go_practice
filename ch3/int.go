//
// main.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  var i int = 100

  // 型が違うとエラー
  var i2 int64 = 200
  // fmt.Println(i + i2)
  fmt.Println(i + 100)

  // 型変換
  // 書式指定子を使って型を表示
  fmt.Printf("%T\n", i2)
  fmt.Printf("%T\n", int32(i2))
  fmt.Println(int64(i) + i2)

  // 浮動小数点
}
