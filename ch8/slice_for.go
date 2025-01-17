//
// slice_for.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  sl := []string{"A", "B", "C"}
  fmt.Println(sl)

  // for i, v := range sl {
  //   fmt.Println(i, v)
  // }

  for i := 0; i < len(sl); i++ {
    fmt.Println(i, sl[i])
  }
}
