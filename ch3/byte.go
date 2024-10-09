//
// byte.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func main(){
  // カギカッコに注意
  byteA := []byte{72, 73}
  fmt.Println(byteA)

  fmt.Println(string(byteA))

  // 文字列をbyteに変換
  // 丸括弧に注意
  c := []byte("HI")
  fmt.Println(c)
  fmt.Println(string(c))
}
