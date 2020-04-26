package main

import "fmt"

func main(){

  jb:=[]string{"James", "Bond", "Shaken, not stirred"}
  mp:=[]string{"Miss", "Moneypenny", "Helloooooo, James."}
  mult_dim :=[][]string{jb,mp}

  for _,v:=range mult_dim{
    for _,v:=range v{
      fmt.Printf("%s\n",v)
    }
  }
}
