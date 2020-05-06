package main

import(
  "os"
  "text/template"
)

type person struct{
  Name string
  Age int
}

func main(){
  p1:=person{"Vipin",25}
  tmp1, err:= template.New("test").Parse("{{.Name}} is {{.Age}} years old")
  if err!=nil{panic(err)}
  err = tmp1.Execute(os.Stdout,p1)
  if err!=nil{panic(err)}

}
