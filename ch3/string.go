//
// string.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func main() {
  var s string = "Hello, World"
  fmt.Println(s)
  fmt.Printf("%T\n", s)

  // string to number
  var si string = "300"
  fmt.Println(si)
  fmt.Printf("%T\n", si)

  // 文字をそのまま表示
  fmt.Println(
  `test
      test
          test`)

  // エスケープ
  fmt.Println("\"")

  // 文字列の抽出
  // byte型として取得
  fmt.Println(s[0])
  fmt.Println(string(s[0]))
}
