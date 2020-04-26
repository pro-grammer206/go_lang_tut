package main

import (
  "fmt"
  "math"
)


type Rect struct{
  length float64
  breadth float64
}

type Circle struct{
  radius float64
}

type Shape interface{
  Area() float64
  Perimeter() float64
}

func (r Rect)Area()float64{
  return r.length*r.breadth
}

func (r Rect)Perimeter()float64{
  return 2*(r.length+r.breadth)
}

func (c Circle)Area()float64{
  return math.Pi*c.radius*c.radius
}

func (c Circle)Perimeter()float64{
  return 2*math.Pi*c.radius
}

func main(){
  var s Shape = Rect{10,3}


  fmt.Printf("type of s is %T\n",s)
  fmt.Printf("value of s is %v\n",s)
  fmt.Printf("value of s is %0.2f\n\n",s.Area())
  fmt.Println("Now s is changed to type circle")
  s = Circle{10}

  fmt.Printf("type of c is %T\n",s)
  fmt.Printf("value of c is %v\n",s)
  fmt.Printf("value of c is %0.2f\n\n",s.Area())


}
