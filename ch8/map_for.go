//
// map_for.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  m := map[string]int{"apple": 100, "banana": 200, "lemon": 300}

  for k, v := range m {
    fmt.Printf("%s => %d\n", k, v)
  }
}
