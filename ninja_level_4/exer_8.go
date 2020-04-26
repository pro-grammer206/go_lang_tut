package main

import "fmt"

func main(){
  m:=map[string][]string{
    `bond_james`:[]string{ `Shaken, not stirred`, `Martinis`, `Women`},
    `moneypenny_miss`:[]string{ `James Bond`, `Literature`, `Computer Science`},
    `no_dr`:[]string{`Being evil`, `Ice cream`, `Sunsets`},
  }
  for k,_:=range m{
    for i,v:=range m[k]{
      fmt.Printf("%d -->%s\n",i,v)
    }
  }
}
