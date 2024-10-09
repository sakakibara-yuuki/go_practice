//
// main.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import (
  "fmt"
  "time"
  "net/http"
  "log"
  "io/ioutil"
  jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySigningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request){
  validToken, err := GenerateJWT()
  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  client := &http.Client{}
  req, _ := http.NewRequest("GET", "http://localhost:9000", nil)
  req.Header.Set("Token", validToken)
  res, err := client.Do(req)
  if err != nil {
    fmt.Fprintf(w, "Error: %s", err.Error())
  }

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  fmt.Fprintf(w, string(body))
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  log.Fatal(http.ListenAndServe(":9001", nil))
}

func GenerateJWT() (string, error){
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)

  claims["authorized"] = true
  claims["user"] = "Elliot Forbes"
  claims["exp"] = time.Now().Add(time.Hour * 30).Unix()

  tokenString, err := token.SignedString(mySigningKey)
  if err != nil {
    fmt.Errorf("Something Went Wrong: %s", err.Error())
    return "", err
  }
  return tokenString, nil
}

func main() {
  fmt.Println("My Simple JWT")
  handleRequests()
}
