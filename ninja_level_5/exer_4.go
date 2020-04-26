package main

import ("fmt")

func main(){
  s1:=struct{
    name string
    age int
  }{
    name:"Vipin",
    age:25,
  }
  fmt.Println(s1.name,s1.age)
}
