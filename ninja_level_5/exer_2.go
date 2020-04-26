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

  m:= map[string]person{
    "Gilly":p2,
    "Doe":p1,
  }
  fmt.Println("Favourite flavours of Jill")
  for _,v:=range m["Gilly"].fav_flavours{
    fmt.Println(v)
  }
  fmt.Println("Favourite flavours of John")
  for _,v:=range m["Doe"].fav_flavours{
    fmt.Println(v)
  }
}
