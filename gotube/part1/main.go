//
// main.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type Article struct {
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
  articles := Articles{
    Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
  }
  fmt.Println("Endpoint Hit: All Articles Endpoint")
  json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {

  myRouter := mux.NewRouter().StrictSlash(true)

  myRouter.HandleFunc("/", homePage)
  myRouter.HandleFunc("/articles", allArticles)
  log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
  handleRequests()
}
