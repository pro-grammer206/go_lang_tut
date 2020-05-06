package main

import (
  "net/http"
  "html/template"
  "os"
)

type Address struct{
  Street string
  Country string
}
type person struct{
  Name string
  Salary int
  Address
}



var tpl *template.Template

func init(){
  tpl = template.Must()
}
