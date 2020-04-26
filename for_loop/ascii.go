package main

import (
  "fmt"
)

func main(){
  for i:=65;i<122;i++{
    fmt.Printf("%#U\t",i)
  }
}
