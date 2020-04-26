package main

import ("fmt")


func concat(s... string)string{
  phrase:=""
  for _,v:=range s{
    phrase+=v
  }
  return phrase

}

func main(){
  fmt.Println(concat("hello"),"friend")

}
