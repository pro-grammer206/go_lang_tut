package main

import (
  "fmt"
  "sort"
)

type Person struct{
  First string
  Age int
}

type ByAge []Person
type ByName []Person

func(a ByAge) Len()int   { return len(a)}
func(a ByAge) Swap(i,j int)   {a[i],a[j]=a[j],a[i]}
func(a ByAge) Less(i,j int)bool   {return a[i].Age < a[j].Age}

func(a ByName) Len()int   { return len(a)}
func(a ByName) Swap(i,j int)   {a[i],a[j]=a[j],a[i]}
func(a ByName) Less(i,j int)bool   {return a[i].First < a[j].First}

func main(){
  p1:=Person{"James",25}
  p2:=Person{"John",32}
  p3:=Person{"Dennis",2}
  p4:=Person{"Jacob",23}
  p5:=Person{"Jennis",2}

  people:=[]Person{p1,p2,p3,p4,p5}

  fmt.Println(people)
  fmt.Println("After sorting by age we get")
  sort.Sort(ByAge(people))
  fmt.Println(people,"\n")

  fmt.Println("After sorting by name we get")
  sort.Sort(ByName(people))
  fmt.Println(people)

}
