package main

import(
  "fmt"
  "encoding/json"
)

type person struct{
  Name string
  Address string
  Age int
}

func main(){

  p1:=person{
    Name:"Vipin",
    Age:25,
    Address:"World",
  }

  p2:=person{
    Name:"Hulk",
    Age:1000000,
    Address:"Universe",
  }

  people:=[]person{p1,p2}

  b,_:=json.Marshal(people)
  fmt.Println(string(b))

}
