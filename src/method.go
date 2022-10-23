package main

import "fmt"

type point struct {
  name string
  point  int
}

func (r point) q()  (string,int){ //func 리시버(변수명  구조체명)
  return "당신의 포인트는", r.point
}

func main() {
  p := point{name: "seungho",point: 1387}
  fmt.Println(  p.q() )
}