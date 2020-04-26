package main

import(
  "fmt"
)

type person struct{
  first_name string
  last_name string
  fav_flavours []string
}

func main(){
  p1:=person{
    first_name:"John",
    last_name:"Doe",
    fav_flavours:[]string{"Chocolate","HazzleNut","Strawberry"},

  }

  p2:=person{
    first_name:"Jill",
    last_name:"Gilly",
    fav_flavours:[]string{"Dark Chocolate","Vanilla","Butter Scotch"},

  }

  for _,v:=range p1.fav_flavours{
    fmt.Println(v)
  }

  for _,v:=range p2.fav_flavours{
    fmt.Println(v)
  }
}
