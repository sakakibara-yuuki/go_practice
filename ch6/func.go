//
// func.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

// 引数の型を省略できる
func Plus(x, y int) int {
  return x + y
}

// 複数の戻り値を返す
func Div(x, y int) (int, int) {
  q := x / y
  r := x % y
  return q, r
}

// 返り値に名前をつけるとreturnに変数を指定しなくても返せる
func Double(price int) (result int) {
  result = price * 2
  return
}

//帰り値の無い関数
func Noreturn() {
  fmt.Println("No Return")
  return
}

func main() {
  i := Plus(1, 2)
  fmt.Println(i)

  i2, i3 := Div(9, 3)
  fmt.Println(i2, i3)

  // 戻り値の一部を無視する
  // i2, _ := Div(9, 3)
  // fmt.Println(i2)

  i4 := Double(1000)
  fmt.Println(i4)

  Noreturn()
}
