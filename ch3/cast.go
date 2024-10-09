//
// cast.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"
// import "strconv"

func main() {
  /*
  var i int = 1
  fl64 := float64(i)
  fmt.Println(fl64)
  fmt.Printf("i = %T\n", i)
  fmt.Printf("fl64 = %T\n", fl64)

  i2 := int(fl64)
  fmt.Printf("i2 = %T\n", i2)

  // 小数点切り捨て
  fl := 10.5
  i3 := int(fl)
  fmt.Printf("i3 = %T\n", i3)
  fmt.Println(i3)
  */

  var s string = "100"
  fmt.Printf("s = %T\n", s)

  // _で破棄することでunusedエラーを回避
  // i, _ := strconv.Atoi(s)
  // fmt.Println(i)
  // fmt.Printf("i = %T\n", i)
  //
  // var i2 int = 200
  // s2 := strconv.Itoa(i2)
  // fmt.Println(s2)
  // fmt.Printf("s2 = %T\n", s2)

  var h string = "Hello, World"
  b := []byte(h)
  fmt.Println(b)

  h2 := string(b)
  fmt.Println(h2)
  
}
