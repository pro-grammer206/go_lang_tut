package main

import (
  "fmt"
)

func main(){
a
:= increementor()
b := increementor()
fmt.Println(a())
fmt.Println(a())
fmt.Println(a())

fmt.Println(b())
fmt.Println(b())
fmt.Println(b())


}

func increementor()func() int{
  var x int
  return func() int{
    x++
    return x
  }

}
