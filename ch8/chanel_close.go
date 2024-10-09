//
// chanel_close.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"
import "time"

func reciever(name string, ch chan int) {
  for {
    i, ok := <-ch
    // okがfalseなら空かつcloseなら終了
    if !ok {
      break
    }
    fmt.Println(name, i)
  }
  fmt.Println(name + "END")
}

func main() {
  ch1 := make(chan int, 2)
  //
  // // deleteとは違う
  // // close(ch1)
  // // closeされたchanelには送ることはできないが, 
  // ch1 <- 1
  // // 取り出すことができる
  // // fmt.Println(<-ch1)
  // close(ch1)
  //
  // // chanelがopenかどうかがわかる
  // // 厳密にはchannelが空でかつcloseならfalse
  // i, ok := <-ch1
  // fmt.Println(i, ok)

  go reciever("1.goroutin", ch1)
  go reciever("2.goroutin", ch1)
  go reciever("3.goroutin", ch1)

  i := 0
  for i < 100 {
    ch1 <- i
    i++
  }
  close(ch1)
  time.Sleep(3 * time.Second)
}
