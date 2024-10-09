//
// array.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func main() {
  // 配列は要素数を指定して宣言する
  // 変更できない
  // 変更する場合はスライスを使う
  var arr1 [3]int
  fmt.Println(arr1)
  fmt.Printf("%T\n", arr1)

  // 初期値を指定する
  // 配列の初期化には{}を使う
  var arr2 = [3]string{"A", "B"}
  // 最後にはカラ文字が入る
  fmt.Println(arr2)

  // 暗黙的型宣言 ?
  arr3 := [3]int{1, 2, 3}
  fmt.Println(arr3)

  // 要素数を省略
  arr4 := [...]string{"C", "D"}
  fmt.Println(arr4)
  fmt.Printf("%T\n", arr4)

  // 配列の要素を取得
  fmt.Println(arr2[0])
  fmt.Println(arr2[1])
  fmt.Println(arr2[2]) // 空文字
  // fmt.Println(arr2[-1]) // エラー
  fmt.Println(arr2[2-1])

  arr2[2] = "C"
  fmt.Println(arr2)

  // 配列のコピーでは要素数が同じでないとコピーできない
  // var arr5 [4]int
  // arr5 = arr1
  // fmt.Println(arr5)

  fmt.Println(len(arr1))
}
