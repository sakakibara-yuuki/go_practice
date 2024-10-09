//
// user.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import (
  "fmt"
  "net/http"
  "encoding/json"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "github.com/gorilla/mux"
)

var db *gorm.DB
var err error

type User struct {
  gorm.Model
  Name string
  Email string
}

func InitialMigration() {
  db, err = gorm.Open(sqlite.Open("test.db"))
  if err != nil {
    fmt.Println(err.Error())
    panic("failed to connect database")
  }

  db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
  db, err = gorm.Open(sqlite.Open("test.db"))
  if err != nil {
    panic("Could not connetct to the database")
  }

  var users []User
  db.Find(&users)
  json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
  db, err = gorm.Open(sqlite.Open("test.db"))
  if err != nil {
    panic("Could not connetct to the database")
  }

  vars := mux.Vars(r)
  name := vars["name"]
  email := vars["email"]
  db.Create(&User{Name: name, Email: email})

  fmt.Fprintf(w, "New User Successfully Created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  db, err = gorm.Open(sqlite.Open("test.db"))
  if err != nil {
    panic("Could not connetct to the database")
  }

  vars := mux.Vars(r)
  name := vars["name"]

  var user User
  db.Where("name = ?", name).Find(&user)
  db.Delete(&user)

  fmt.Fprintf(w, "User Successfully Deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  db, err = gorm.Open(sqlite.Open("test.db"))
  if err != nil {
    panic("Could not connetct to the database")
  }

  vars := mux.Vars(r)
  name := vars["name"]
  email := vars["email"]

  var user User
  db.Where("name = ?", name).Find(&user)
  user.Email = email
  db.Save(&user)

  fmt.Fprintf(w, "User Successfully Update")
}
