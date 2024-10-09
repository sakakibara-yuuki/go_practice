//
// defer.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"
import "os"

// 関数が終了する際にdeferが実行される
func TestDefer() {
  defer fmt.Println("END")
  fmt.Println("START")
}

func RunDefer() {
  defer fmt.Println("1")
  defer fmt.Println("2")
  defer fmt.Println("3")
}

func main() {
  TestDefer()

  // 無名関数で処理をまとめる
  // defer func() {
  //   fmt.Println("1")
  //   fmt.Println("2")
  //   fmt.Println("3")
  // }()

  RunDefer()

  file, err := os.Create("test.txt")
  if err != nil {
    fmt.Println("Failed to create a file")
  }
  defer file.Close()

  file.Write([]byte("Hello, World!"))
}
