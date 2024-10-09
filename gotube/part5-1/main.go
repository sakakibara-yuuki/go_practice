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

type SensorReading struct {
  Name string `json:"name"`
  Capacity int `json:"capacity"`
  Time string `json:"time"`
  Information Info `json:"info"`
}

type Info struct {
  Description string `json:"desc"`
}

func main() {
  jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2020-01-21T19:07:28Z", "info": {"desc": "a sensor reading"}}`

  var reading SensorReading
  err := json.Unmarshal([]byte(jsonString), &reading)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Printf("%+v\n", reading)
}
