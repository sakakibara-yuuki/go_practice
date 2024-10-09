//
// error.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"
import "strconv"

func main() {

  var s string = 'A'

  i, err := strconv.Atoi("45")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(i)
  }
}
