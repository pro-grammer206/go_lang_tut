package main

import(
  "database/sql"
  "fmt"
  _"github.com/lib/pq"

)

func main(){
  fmt.Println(sql.Drivers())
}
