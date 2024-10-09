//
// check.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func call(a *int) {
  fmt.Println(a + 10);
}

func main() {
  var n int = 100;
  var p *int = &n;
  fmt.Println(p);
}
