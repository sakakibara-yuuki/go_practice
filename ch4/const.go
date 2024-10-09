//
// const.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

// 頭文字を大文字にすると他のパッケージからも参照できる

// 定数
const Pi = 3.14

const (
  URL = "http://example.com"
  SiteName = "test"
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

const (
  A = 1
  B
  C
  D= "A"
  E
  F
)

const (
  c0 = iota
  c1
  c2
)

func main() {
  fmt.Println(Pi)
  // 一度定義した定数は変更できない
  // Pi = 3
  // fmt.Println(Pi)

  // 複数の定数をまとめて定義する
  fmt.Println(URL, SiteName)

  // 値の省略
  fmt.Println(A, B, C, D, E, F)
  // 値を正楽すると, その値が続いているものを引き継ぐ

  // 巨大な数を扱う
  // 通常はオーバーフローする
  fmt.Println(Big - 1)

  // iota
  fmt.Println(c0, c1, c2)
  // 連番を生成する
}
