package main

import "fmt"

func main(){
  a:=[]int{0,1,2,3,4,5,6,7,8,9}
  for i,v:=range a{
    fmt.Printf("%d\t%d\n",i,v)
  }
  fmt.Printf("type of a is %T",a)

}
