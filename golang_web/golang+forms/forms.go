package main

import (
  "html/template"
  "log"
  "net/http"
)

var tpl *template.Template

type Fullname struct{
  First string
  Last string
}

func init(){
  tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
  http.HandleFunc("/",index)
  http.HandleFunc("/process",process)
  http.ListenAndServe(":5000",nil)
}

func index(w http.ResponseWriter, r *http.Request){
  err:=tpl.ExecuteTemplate(w,"form.gohtml",nil)
  if err!=nil{
    log.Fatal(err)
  }
}

func process(w http.ResponseWriter, r*http.Request){
  if r.Method=="POST"{

    // fname:=r.FormValue("fname")
    // lname:=r.FormValue("lname")

    f:=Fullname{
      First:r.FormValue("fname"),
      Last:r.FormValue("lname"),
    }
    tpl.ExecuteTemplate(w,"display.gohtml",f)
  }else{
    http.Redirect(w,r,"/",http.StatusSeeOther)
    return
  }
}
