package main

import (
  "fmt"
)

type person struct{
  first string
  last string
  age int
}

func (p person) speak(){
  fmt.Println("My name is",p.first,p.last,",I am",p.age,"years old")
}

func main(){
  p:=person{"Hulk","Smash",100000}
  p.speak()

}
