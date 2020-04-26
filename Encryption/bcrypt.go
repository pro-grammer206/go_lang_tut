package main

import(
  "fmt"
  "golang.org/x/crypto/bcrypt"
)

func main(){
  s:="password123"
  bs,err:=bcrypt.GenerateFromPassword([]byte(s),bcrypt.MinCost)
  if err!=nil{
    fmt.Println(err)
  }
  fmt.Println(s)
  fmt.Println(string(bs))


  loginpwd:=`passfword123`
  err = bcrypt.CompareHashAndPassword(bs,[]byte(loginpwd))
  if err!=nil{
    fmt.Println("You cannot login")
    return
  }
  fmt.Println("You are logged in")
}
