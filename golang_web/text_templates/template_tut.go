package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	name := os.Args[1]
	str := `
<!DOCTYPE html>
 <html lang="en">
    <head>
        <title>Hello Go Language</title>
      </head>
        <body>
          <h1>` + name + `</h1>
        </body>
  </html>`

	nf, err := os.Create("Index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}
