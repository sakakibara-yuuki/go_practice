//
// chanel.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
  var ch1 chan int
  //受信専用のチャネル
  // var ch2 <-chan int
  //送信専用のチャネル
  // var ch3 chan<- int

  ch1 = make(chan int)
  ch2 := make(chan int)
  fmt.Println(cap(ch1))
  fmt.Println(cap(ch2))

  ch3 := make(chan int, 5)
  fmt.Println(cap(ch3))

  // 入れたぶんだけlenが増える
  ch3 <- 1
  fmt.Println(len(ch3))
  ch3 <- 2
  ch3 <- 3
  // fmt.Println(len(ch3))
  fmt.Println("len", len(ch3))

  // queになってる
  i := <- ch3
  fmt.Println(i)
  fmt.Println("len", len(ch3))
  i2 := <- ch3
  fmt.Println(i2)
  fmt.Println("len", len(ch3))

  fmt.Println(<-ch3)
  fmt.Println("len", len(ch3))

  // buffer以上を送るとdeadlockになる
  ch3 <- 1
  // fmt.Println(<-ch3)
  ch3 <- 2
  ch3 <- 3
  ch3 <- 4
  ch3 <- 5
  ch3 <- 6

}
