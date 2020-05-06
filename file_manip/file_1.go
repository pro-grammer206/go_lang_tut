package main

import (
  "fmt"
  "os"
  "bufio"
)

func main(){
  // f,err := os.Create("data.txt")
  // if err!=nil{
  //   panic(err)
  // }
  f,err:=os.Open("data.txt")
  defer f.Close()

  s:=bufio.NewScanner(f)
  for s.Scan(){
    fmt.Println(s.Text())
  }
  if err=s.Err();err!=nil{
    fmt.Fprintln(os.Stderr,"reading standard input:",err)
  }

}
