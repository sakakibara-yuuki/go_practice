//
// main.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import (
  "fmt"
  "net/http"
  "log"
  "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    if r.Header["Token"] != nil {
      token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf("There was an error")
        }
        return mySigningKey, nil
      })

      if err != nil {
        fmt.Fprintf(w, err.Error())
      }

      if token.Valid {
        endpoint(w, r)
      }

    } else {
      fmt.Fprintf(w, "Not Authorized")
    }

  })
}

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Super Secret Information")
  fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
  // http.HandleFunc("/", homePage)
  http.Handle("/", isAuthorized(homePage))
  log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
  fmt.Println("My Simple Server")
  handleRequests()
}
