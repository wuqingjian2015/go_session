package main 

import (
	"fmt"
	"os"
	"text/template"
)

func main(){
	t := template.New("test")
	t=template.Must(t.ParseFiles("./header.tmpl", "./content.tmpl", "./footer.tmpl"))
	var v = 1
	t.ExecuteTemplate(os.Stdout, "header", v)
	fmt.Println()
	t.ExecuteTemplate(os.Stdout, "content", v)
	fmt.Println()
	t.ExecuteTemplate(os.Stdout, "footer", v)
	// fmt.Println()
	// t.Execute(os.Stdout, nil)
	tOk := template.New("first")
	template.Must(tOk.Parse("some static text /*commented out*/"))
	fmt.Println("the first one parsed ok")

	template.Must(template.New("second").Parse("some static text {{.Name}}"))
	fmt.Println("the second one parsed ok")

	template.Must(template.New("third").Parse("some static text {{.Name}"))
	
}