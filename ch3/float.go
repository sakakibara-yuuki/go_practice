//
// float.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func main(){
  var fl64 float64 = 2.4
  fmt.Println(fl64)

  fl := 3.2
  fmt.Println(fl + fl64)
  fmt.Printf("%T, %T\n", fl, fl64)

  var fl32 float32 = 1.2
  fmt.Printf("%T\n", fl32)

  // 型変換
  fmt.Printf("%T\n", float64(fl32))

  // floatにならない
  zero := 0.0
  pinf := 1.0 / zero
  fmt.Println(pinf)

  ninf := -1.0 / zero
  fmt.Println(ninf)

  nan := zero / zero
  fmt.Println(nan)
}
