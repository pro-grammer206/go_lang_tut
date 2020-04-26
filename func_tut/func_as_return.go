package main

import "fmt"

func main(){

  pancakeMaker:=makePancake()
  fmt.Println(pancakeMaker())
  fmt.Printf("type of pancakeMaker is %T",pancakeMaker)

}

func makePancake()func()string{
  return func()string{
    return "Ta da pancake is ready"
  }
}
