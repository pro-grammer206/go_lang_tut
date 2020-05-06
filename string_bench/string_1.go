package main

import(
  "fmt"
  "reflect"
)
type IPAddr []byte
func main(){
  a:="this is a string text"
  b:=a+"lalaalal"

  fmt.Println(reflect.TypeOf(a[1]))
  fmt.Println(a[1])
  fmt.Printf("%s",b)
}
