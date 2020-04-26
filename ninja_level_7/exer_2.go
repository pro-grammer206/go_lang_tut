package main

import (
  "fmt"
)
type person struct{
  name string
  address string
  age int
}

func changeMe(p *person){
  (*p).name = "Hulk"
  (*p).address = "Universe"

}
func main(){
  p1:=person{
    name:"Vipin",
    address:"India",
    age:25,
  }
  changeMe(&p1)
  fmt.Println(p1)

}
