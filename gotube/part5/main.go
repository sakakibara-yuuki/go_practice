//
// main.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import (
  "fmt"
  "encoding/json"
)

type Book struct {
  Title string `json:"title"`
  Author Author `json:"author"`
}

type Author struct {
  Name string `json:"name"`
  Age int `json:"age"`
  Developer bool `json:"developer"`
}

func main() {
  fmt.Println("Hello, World!")

  book := Book{Title: "Learning Concurrency in Python", Author: Author{Name: "Elliot Forbes", Age: 25, Developer: true}}
  fmt.Println(book)
  fmt.Printf("%+v\n", book)

  byteArray, err := json.MarshalIndent(book, "", "  ")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(byteArray))

}
