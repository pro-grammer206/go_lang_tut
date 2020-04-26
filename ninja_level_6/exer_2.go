package main

import (
  "fmt"
)

func main(){
  fmt.Println("The sum of numbers 1 to 9 is",foo([]int{1,2,3,4,5,6,7,8,9}...))

}

func foo(x...int)int{
  sum:=0
  for _,v:=range x{
    sum+=v
  }
  return sum
}
