package main

import(
  "fmt"
)

func main(){
  i:=1
  fmt.Printf("Before function call value of i is %d\n",i)
  foo(&i)
  fmt.Println("After function call value if i is",i)

}

func foo(j *int){
  fmt.Println("Value of j is",*j,"\n")
  fmt.Printf("type of j is %T\n",j)
  *j = 42

}
