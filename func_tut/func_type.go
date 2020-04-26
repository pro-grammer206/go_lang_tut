package main

import "fmt"

func main(){
  toDoList:=func(){
    fmt.Println(`1.Intellectual
2.Muscular
3.Repeat`)
  }
  toDoList()
  fmt.Printf("Type of todolist is %T\n",toDoList)

}
