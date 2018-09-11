package main

import (
	"os"
	"text/template"
)
func main(){
	t := template.New("test")
	t = template.Must(t.Parse("空 pipeline if demo: {{ if `` }} 不会输出 {{end}}\n"))
	t.Execute(os.Stdout, nil)

	t1 := template.New("test with value")
	t1 = template.Must(t.Parse("non-empty pipeline if demo: {{ if `something` }} some content {{end}}\n"))

	t1.Execute(os.Stdout, nil)

	t2 := template.New("test if else with value")
	t2 = template.Must(t.Parse("non-empty pipeline if-else demo: {{if `some` }} empty content {{ else }} else content {{end}}\n"))
	t2.Execute(os.Stdout, nil)




}