//
// label_for.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

func main() {
// Loop:
//   for {
//     for {
//       for {
//         fmt.Println("Looping")
//         break Loop
//       }
//       fmt.Println("This line will never print")
//     }
//     fmt.Println("This line will never print")
//   }
//   fmt.Println("End")

Loop:
  for i := 0; i < 3; i++ {
    for j:=1; j < 3; j++ {
      if j > 1 {
        continue Loop
      }
      fmt.Println(i, j, i*j)
    }
    fmt.Println("処理をしない")
  }

}
