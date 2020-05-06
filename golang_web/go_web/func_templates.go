package main

import (
  "text/template"
  "math"
  "log"
  "os"
  "time"
)
var tpl *template.Template

type Mathtime struct{
  Num int
  Curr_time time.Time

}

func init(){
  tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

}
func double(x int)int{
  return x + x
}

func square(x int)float64{
  return math.Pow(float64(x),2)
}

func sqRoot(x float64)float64{
  return math.Sqrt(x)
}

func dayMonthYear(t time.Time)string{
  return t.Format("02-01-2006")
}

var fm = template.FuncMap{
  "fdbl":double,
  "fsq":square,
  "fsqrt":sqRoot,
  "fdmy":dayMonthYear,
}

func main(){
  mt:=Mathtime{
    Num:9,
    Curr_time:time.Now(),
  }
  err:= tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",mt)
  if err!=nil{
    log.Fatal(err)
  }
}
