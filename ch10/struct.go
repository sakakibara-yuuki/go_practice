//
// struct.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"


type User struct {
  Name string
  Age int
  // X, Y int
}

func UpdateUser(user User) {
  user.Name = "A"
  user.Age = 100
}

func UpdateUser2(user *User) {
  user.Name = "A"
  user.Age = 100
}


func main() {
  var user1 User
  fmt.Println(user1)
  user1.Name = "sakakibara"
  user1.Age = 10
  fmt.Println(user1)

  user2 := User{}
  user2.Name = "sakakibara2"
  fmt.Println(user2)

  user3 := User{"sakakibara", 30}
  fmt.Println(user3)

  user4 := User{Age: 40, Name: "sakakibara"}
  fmt.Println(user4)

  user6 := User{Name: "user6"}
  fmt.Println(user6)

  user5 := new(User)
  user5.Name = "sak2akibara"
  user5.Age = 1000
  fmt.Println(*user5)

  // こちらがよく使われる
  // 関数の引数など
  user7 := &User{Name: "sakakibara", Age: 30}
  fmt.Println(user7)

  UpdateUser(user1)
  UpdateUser2(user7)
  fmt.Println(user1)
  fmt.Println(user7)
}
