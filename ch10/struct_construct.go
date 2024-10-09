//
// struct_construct.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

type User struct {
  Name string
  Age int
}

// 型としては*はポインタや参照を表すが
// 演算子としてはデリファレンスを表す
// 演算子としての&はアドレスを取得する
// factory のようなものを作る
func NewUser(Name string, Age int) *User {
  return &User{Name, Age}
}

func main() {
  user1 := NewUser("user1", 20)
  fmt.Println(user1)
  fmt.Println(*user1)
}
