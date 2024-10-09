//
// main.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

var i5 int = 100

func outer() {
  var s4 string = "outer"
  fmt.Println(s4)
}

func main() {
  // 明示的な定義
  // var 変数名 型 = 値
  var i int = 1000
  fmt.Println(i)

  // 文字列
  var s string = "Hello, World!"
  fmt.Println(s)

  // 複数
  var t, f bool = true, false
  fmt.Println(t, f)

  // 異なる型
  var (
    i2 int = 200
    s2 string = "Golang"
  )
  fmt.Println(i2, s2)

  // 初期値なし
  var i3 int
  var s3 string
  fmt.Println(i3, s3)
  // このような場合, int は 0 string は "" で初期化される""(空文字)なので表示されない

  // 代入
  i3 = 300
  s3 = "Go"
  fmt.Println(i3, s3)

  i = 150
  fmt.Println(i)

  // 暗黙的な定義
  i4 := 400
  fmt.Println(i4)
  // varも必要ないし型も必要ない.

  // これは再代入
  i4 = 450
  fmt.Println(i4)

  // 暗黙的な定義は再定義はできない
  // i4 := 500
  // fmt.Println(i4)
  // 異なる型で代入はできない.

  // 暗黙的な型定義は関数内でしか使えない
  fmt.Println(i5)

  // 一般的には型指定を行う

  outer()

  // 関数内の変数は関数内でしか使えない
  // fmt.Println(s4)

  // 使われない変数はエラーになる
  // var s5 string = "Not Use"
}
