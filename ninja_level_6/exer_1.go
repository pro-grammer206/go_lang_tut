package main

import (
  "fmt"
)

func main(){
  fmt.Println(foo())
  name,age:=bar()
  fmt.Println(name,age)
  fmt.Println(bar())

}

func foo()int{
  return 20
}
func bar()(string,int){
  return "Vipin",20
}
