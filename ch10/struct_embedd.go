//
// struct_embedd.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

type T struct {
  // User User
  User
}

type User struct {
  Name string
  Age int
}

func (u *User) SetName() {
  u.Name = "A"
}

func main() {
  t := T{
    User: User{
      Name: "sakakibara",
      Age: 100,
    },
  }

  fmt.Println(t)
  fmt.Println(t.User)
  fmt.Println(t.User.Name)
  fmt.Println(t.User.Age)
  // 省略するとT.Nameのように間を省略できる
  fmt.Println(t.Name)

  t.User.SetName()
  fmt.Println(t)
}
