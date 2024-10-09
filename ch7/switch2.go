//
// switch2.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func anything(a interface{}) {
  fmt.Println(a)
  switch v := a.(type) {
    case int:
      fmt.Println(v + 100)
    case string:
      fmt.Println(v + "hoge")
    default:
      fmt.Println("I don't know")
  }
}

func main() {
  anything("aa")
  anything(1)

  var x interface{} = 3
  i := x.(int)
  fmt.Println(i + 2)
  // fmt.Println(x + 2)

  // 異なる型に変換するとエラーになる
  // f := x.(float64)
  // fmt.Println(f + 0.1)

  // 型アサーションの結果を使わない場合
  f, isFloat64 := x.(float64)
  fmt.Println(f, isFloat64)

  if x == nil {
    fmt.Println("None")
  } else if i, isInt := x.(int); isInt {
    fmt.Println(i, "x is int")
  } else if s, isString := x.(string); isString {
    fmt.Println(s, "x is string")
  } else {
    fmt.Println("I don't know")
  }

  switch x.(type) {
    case int:
      fmt.Println("int")
    case string:
      fmt.Println("string")
    default:
      fmt.Println("I don't know")
  }

  switch v := x.(type) {
    case bool:
      fmt.Println(v, "bool")
    case int:
      fmt.Println(v, "int")
    case string:
      fmt.Println(v, "string")
    default:
      fmt.Println("I don't know")
  }
}
