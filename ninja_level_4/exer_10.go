package main

import "fmt"

func main(){
  m:=map[string][]string{
    `bond_james`:[]string{ `Shaken, not stirred`, `Martinis`, `Women`},
    `moneypenny_miss`:[]string{ `James Bond`, `Literature`, `Computer Science`},
    `no_dr`:[]string{`Being evil`, `Ice cream`, `Sunsets`},
  }
  m[`kumar_vipin`]=[]string{`Intellectual`,`Muscular`,`rich`}
  delete(m,`moneypenny_miss`)
  for k,_:=range m{
    for i,v:=range m[k]{
      fmt.Printf("%d -->%s\n",i,v)
    }
  }
}
