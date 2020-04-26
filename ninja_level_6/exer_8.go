package main

import (
  "fmt"
  "math"
)

func main(){
  square:=squared(2)
  fmt.Println(square())
  fmt.Println(square())
  fmt.Println(square())


}

func squared(a float64)func()float64{
  a  := 1
  return func()float64{
    a = math.Pow(a,2)
    return a
  }
}
