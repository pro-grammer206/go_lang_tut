package main

import (
  "fmt"
  "sort"
  "os"
)

func main(){
  nums:= []int{12,89,2,96,25,45,1,9,4,5,6,4,200}
  fmt.Println("Before sorting nums array was",nums)
  sort.Ints(nums)
  fmt.Println("After sorting nums arrays is",nums)
  os.Stdout.WriteString("This is a type of output to console")
}
