//
// chan_go.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"
import "time"

func reciver(c chan int) {
  // awaitしてるの?
  // なぜこれが動く?
  for {
    i := <-c
    fmt.Println(i)
  }
}

func main() {
  ch1 := make(chan int)
  ch2 := make(chan int)
  // fmt.Println(<-ch1)

  go reciver(ch1)
  go reciver(ch2)

  i := 0
  for i < 100 {
    ch1 <- i
    ch2 <- i
    time.Sleep(50 * time.Millisecond)
    i++
  }
}
