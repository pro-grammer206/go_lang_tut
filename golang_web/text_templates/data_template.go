package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("life.html"))
}

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Muhammad", "Jesus"}
	err := tpl.ExecuteTemplate(os.Stdout, "life.html", sages)
	if err != nil {
		log.Fatalln(err)
	}

}
