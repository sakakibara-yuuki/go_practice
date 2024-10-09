//
// struct_slice.go
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

type Users []*User

// type Users struct {
//   Users []*User
// }

// &は呼び出しにしか使わない
func main() {
  user1 := User{"user1", 10}
  user2 := User{"user2", 20}
  user3 := User{"user3", 30}
  user4 := User{"user4", 40}

  // users := Users{&user1, &user2, &user3, &user4}
  users := Users{}
  users = append(users, &user1)
  users = append(users, &user2)
  users = append(users, &user3, &user4)

  for _, u := range users {
    fmt.Println(*u)
  }

  // スライスなのでmakeが使える
  // users2 := make([]*User, 0)
  users2 := make(Users, 0)
  users2 = append(users2, &user1, &user2)

  for _, u := range users2 {
    fmt.Println(*u)
  }
}
