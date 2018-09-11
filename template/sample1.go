package main 
import (
	"os"
	"strings"
	"fmt"
	"text/template"
)
type Friend struct {
	Fname string 
}

type Person struct {
	UserName string 
	Emails [] string 
	Friends []*Friend 
}
func DealWithEmail(args ...interface{}) string{
	var s string 
	ok := false 
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	
	if !ok {
		s = fmt.Sprint(args...)
	}
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s 
	}

	return substrs[0] + " at " + substrs[1]
}
func main(){
	f1 := Friend{Fname:"qingjian.wu"}
	f2 := Friend{Fname: "Ken.Ng"}
	t := template.New("field sample")
	t = t.Funcs(template.FuncMap{"emailDeal":DealWithEmail})
	t, _ = t.Parse(`Hello {{.UserName}}!
				{{range .Emails}}
					an emails {{.|emailDeal}}
				{{end}}
				{{with .Friends}}
					{{range .}}
						my friend name is {{.Fname}}
					{{end}}
				{{end}}
					`)
	p := Person{UserName:"kenWu", Emails: []string {"ken@qq.com", "x@icloud.com"},
		Friends:[]*Friend{&f1,&f2}}
	t.Execute(os.Stdout, p)
}