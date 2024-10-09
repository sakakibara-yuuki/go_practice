//
// custom_error.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

// こういうものが背後に定義されていると考える.
// type error interface {
//   Error() string
// }

type MyError struct {
  Msg string
  Code int
}

func (e *MyError) Error() string {
  return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func RaiseError() error {
  return &MyError{Msg: "Error", Code: 123}
}

func main() {
  err := RaiseError()
  fmt.Println(err.Error())

  e, ok := err.(*MyError)
  if ok {
    fmt.Println(e.Code)
  }
}
