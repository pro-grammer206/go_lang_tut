package main

import(
  "fmt"
  "runtime"
  "sync"
)

func main(){
  fmt.Println("CPUs:",runtime.NumCPU())
  fmt.Println("Goroutines:",runtime.NumGoroutine())

  counter:=0

  const gs = 100
  var wg sync.WaitGroup
  wg.Add(gs)

  for i:=0;i<gs;i++{
    go func(){
      v:=counter
      v++
      counter = v
    }()
  }
  fmt.Println("Goroutines:",runtime.NumGoroutine())
  fmt.Println("ta da")
}
