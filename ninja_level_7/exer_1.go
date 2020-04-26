package main

import (
  "fmt"
)

func main(){
  var b int
  c:= &b
  *c=300
  a:=100
  fmt.Println("The address of a is %v",&a)
  fmt.Println("The address of a is %p",&a)
  fmt.Println("The address of a is %d",&a)
  fmt.Println("The value of c is %d",a)
  fmt.Println("The address of b is %d",c)
}
