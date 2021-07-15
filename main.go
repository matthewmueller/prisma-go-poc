package main

import (
	_ "embed"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/prisma/prisma-go/dmmf"
)

//go:embed template.gotext
var templateData string

func main() {
	buf, err := ioutil.ReadFile("dmmf.json")
	if err != nil {
		log.Fatal(err)
	}
	var state dmmf.State
	if err := json.Unmarshal(buf, &state); err != nil {
		log.Fatal(err)
	}
	tpl := template.Must(template.New("template.gotext").Parse(templateData))
	if err := tpl.Execute(os.Stdout, &state); err != nil {
		log.Fatal(err)
	}

	// ok, err := json.Marshal(state)
	// fmt.Println(string(ok))
}
