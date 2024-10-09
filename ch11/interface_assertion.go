//
// interface_assertion.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"


type Stringfy interface {
  ToString() string
}

type Person struct {
  Name string
  Age int
}

func (p *Person) ToString() string {
  return fmt.Sprintf("Name=%v Age=%v", p.Name, p.Age)
}

type Car struct {
  Number string
  Model string
}

func (c *Car) ToString() string {
  return fmt.Sprintf("Number=%v Model=%v", c.Number, c.Model)
}

func main() {
  // 異なるデータ型で合っても, 型の共通部分を抽出したインターフェースを定義することで, それらをまとめて扱うことができる
  vs := []Stringfy{
    &Person{Name: "Taro", Age: 21},
    &Car{Number: "XXX-0123", Model: "Model X"},
  }

  for _, v := range vs {
    fmt.Println(v.ToString())
  }
}
