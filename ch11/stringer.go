//
// stringer.go
// Copyright (C) 2024 sakakibara <sakakibara@organon>
//
// Distributed under terms of the MIT license.
//

package main
import "fmt"

// こういうものが背後に定義されていると考える.
// type Stringer interface {
//   String() string
// }
type Point struct {
  A int
  B int
}

func (p *Point) String() string {
  return fmt.Sprintf("A: %d, B: %d", p.A, p.B)
}

func main() {
  p := Point{A: 1, B: 2}
  // p.String() が裏で呼ばれている.
  // これはpythonの__str__メソッドに似ている.
  fmt.Println(p)
}
