package main

import (
  "fmt"
  "encoding/json"
)

type person struct{
  Name string `json:"Name"`
  Address string `json:"Address"`
  Age int `json:"Age"`
}

func main(){
  s:=`[{"Name":"Vipin","Address":"World","Age":25},{"Name":"Hulk","Address":"Universe","Age":1000000}]`
  b:=[]byte(s)
  fmt.Printf("%T\n",s)
  fmt.Printf("%v\n",string(b))

  var people []person

  err:=json.Unmarshal(b,&people)
  if err!=nil{
    fmt.Println(err)
  }

  fmt.Println("all data",people)

  for i,v:=range people{
    fmt.Println("Person Number",i)
    fmt.Println(v.Name,v.Address,v.Age,"\n\n")
  }


}
