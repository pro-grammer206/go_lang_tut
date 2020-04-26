package main

import (
  "fmt"
  "math"
)

type SQUARE struct{
  length float64
}

type CIRCLE struct{
  radius float64
}

func (s SQUARE) area()float64{
  return s.length*s.length
}

func (c CIRCLE) area()float64{
  return math.Pi*c.radius*c.radius
}

type SHAPE interface{
  area()float64
}

func INFO(s SHAPE){
  fmt.Println("The area of %T is %f",s,s.area())
}

func main(){
  var s,c SHAPE
  s=SQUARE{10}
  c=CIRCLE{20}
  INFO(s)
  INFO(c)


}
