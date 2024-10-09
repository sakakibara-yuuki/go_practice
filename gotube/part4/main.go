//
// main.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  name string `json:"name"`
}


func main() {
  fmt.Println("Go MySQL Tutorial")


  db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  results, err := db.Query("SELECT name FROM users")
  if err != nil {
    panic(err.Error())
  }

  for results.Next() {
    var user User
    err = results.Scan(&user.name)
    if err != nil {
      panic(err.Error())
    }
    fmt.Println(user.name)
  }

  // insert, err := db.Query("INSERT INTO users VALUES('test')")
  // if err != nil {
  //   panic(err.Error())
  // }
  // defer insert.Close()

  fmt.Println("Successfully connected to the database")
}
