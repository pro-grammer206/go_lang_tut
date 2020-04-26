package main

import (
  "fmt"
)

type vehicle struct{
  door int
  color string
}
type truck struct{
  vehicle
  fourWheel bool

}
type sedan struct{
  vehicle
  luxury bool
}

func main(){
  t:=truck{
    vehicle:vehicle{
      door:4,
      color:"Blue",
    },
    fourWheel:true,
  }
  s:=sedan{
    vehicle:vehicle{
      door:5,
      color:"Steel Earth",
    },
    luxury:false,
  }
  fmt.Println(t.door,t.color,t.fourWheel)
  fmt.Println(s.door,s.color,s.luxury)
}
