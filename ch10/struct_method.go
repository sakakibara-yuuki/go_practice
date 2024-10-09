//
// struct_method.go
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

// 何じゃこれ
// レシーバーを使ってメソッドを定義する
func (u User) SayName() {
  fmt.Println(u.Name)
}

// これだと更新されない
func (u User) SetName(name string) {
  u.Name = name
}

// Cならばarrow演算子でアクセスできるが、Goはポインタを使う
// 基本的にはポインタを使う
// 呼び出し側はポインタでなくて良い
func (u *User) SetName2(name string) {
  u.Name = name
}

func main() {
  user1 := User{Name: "user1"}
  user1.SayName()

  user1.SetName("user2")
  user1.SayName()

  user1.SetName2("user3")
  user1.SayName()

  // poinntaでも値でもどっちでもいい
  user2 := &User{Name: "user2"}
  user2.SetName2("B")
  user2.SayName()
}
