package main

import (
	"os"
	"text/template"
)

func main() {
	out, _ := os.Create("./out.html")
	defer out.Close()

	out.Truncate(0)
	template.Must(template.New("").ParseGlob("./src/*")).ExecuteTemplate(out, "index.tpl", struct{}{})

	out.Truncate(0)
	template.Must(template.New("").ParseGlob("./src/*")).ExecuteTemplate(out, "index.tpl", struct{}{})

}
