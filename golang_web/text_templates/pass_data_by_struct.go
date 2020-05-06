package main

import (
  "os"
  "log"
  "text/template"
)

var tpl *template.Template

type Sage struct{
  Name string
  Motto string
}

func init(){
  tpl=template.Must(template.ParseFiles("list.html"))
}
func main(){
  buddha:= Sage{
    Name: "Buddha",
    Motto: "The belief of no beliefs",
  }
  err:=tpl.Execute(os.Stdout,buddha)
  if err!=nil{
    log.Fatalln(err)
  }
}
