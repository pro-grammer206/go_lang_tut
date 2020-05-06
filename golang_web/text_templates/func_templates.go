package main

import (
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("list.html"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s

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
}
