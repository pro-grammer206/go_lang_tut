package main

import (
  _"fmt"
  "html/template"
  "net/http"
)
var tpl *template.Template

func init(){
  tpl=template.Must(template.ParseFiles("layout.html"))
}

type Person struct{
  Name string
  Cars []Car
}

type Car struct{
  Brand string
  Cost int
}

func main(){
  http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
    data:=Person{
      Name:"Vipin",
      Cars:[]Car{
        {"Porshe",1000000},
        {"BMW",20000},
        {"Lamborghini",5000000},
      },
    }
    tpl.Execute(w,data)


  })
  http.ListenAndServe(":5000",nil)
}
