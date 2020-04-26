package main

import ("fmt")

func main(){

  fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
  fmt.Println("The sum of even numbers is",evenSum(sum,[]int{1,2,3,4,5,6,7,8,9}))

}

func sum(x...int)int{
  total:=0
  for _,v:=range x{
    total+=v
  }
  return total


}
func evenSum(f func(x...int)int,y[]int)int{
  var xeven []int
  for _,v:= range y{
    if v%2!=0{
    xeven = append(xeven,v)
  }
  }
  total := f(xeven...)
  return total
}
