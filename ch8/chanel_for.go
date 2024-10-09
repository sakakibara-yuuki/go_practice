//
// chanel_for.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  ch1 := make(chan int, 3)
  ch1 <- 1
  ch1 <- 2
  ch1 <- 3
  close(ch1)
  // 空ではないと, deadlockが発生する
  for i := range ch1 {
    fmt.Println(i)
  }
}
