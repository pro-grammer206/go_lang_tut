package main


import (
  "fmt"
)

func main(){

  foo()

  func(x int){
    fmt.Println("This is an annonymous function with single integer arguement of ",x)

  }(10)

  f:=func(x int){
    fmt.Println("My first func expression")
  }(1984)


}

func foo(){
  fmt.Println("This is a normal foo function")
}
