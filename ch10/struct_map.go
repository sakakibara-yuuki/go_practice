//
// struct_map.go
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

func main() {

  m := map[int]User{
    1: {Name: "Alice", Age: 25},
    2: {Name: "Bob", Age: 30},
  }
  fmt.Println(m)

  m2 := map[User]string{
    {"Alice", 25}: "Tokyo",
    {"Bob", 30}: "LA",
  }
  fmt.Println(m2)

  // mapなのでmakeを使って初期化できる.
  m3 := make(map[int]User)
  fmt.Println(m3)
  m3[1] = User{Name: "Alice", Age: 25}
  fmt.Println(m3)

  for _, v := range m {
    fmt.Println(v)
  }
}
