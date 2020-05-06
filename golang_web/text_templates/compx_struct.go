package main

import (
  "os"
  "text/template"
  "log"
)

var tpl *template.Template
type sage struct{
  Name string
  Motto string
  }

type car struct{
  Manufacturer string
  Model string
  Doors  int
}

type items struct{
  Wisdom []sage
  Transport []car
}

func init(){
  tpl = template.Must(template.ParseFiles("list.html"))

}

func main(){

  s:=sage{
    Name:"Buddha",
    Motto:"The belief of no beliefs",
  }
  g:=sage{
    Name:"Gandhi",
    Motto:"Be the change",
  }
  m:=sage{
    Name:"Martin Luther King",
    Motto:"Hatred never ceases with hatred but with love alone is healed",
  }
  f:=car{
    Manufacturer:"Ford",
    Model: "F150",
    Doors: 2,
  }
  c:=car{
    Manufacturer:"Toyota",
    Model:"Corolla",
    Doors:4,
  }

  sages:= []sage{s,g,m}
  cars:=[]car{f,c}

  data:=items{
    Wisdom: sages,
    Transport:cars,
  }

  err:=tpl.Execute(os.Stdout,data)
  if err!=nil{
    log.Fatalln(err)
  }

}
