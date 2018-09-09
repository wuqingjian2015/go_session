package main

import (
	"net/http"
	"html/template"
	"fmt"
	"session"
	"memory"
	"io"
)
var globalSession1 *session.Manager

var pder *memory.Provider

func main() {
	// init()
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/count", count)
	err := http.ListenAndServe(":5000", nil)
	if err != nil  {
		panic("")
	}
	fmt.Println("app:")
	fmt.Println(globalSession1)

}
func index(w http.ResponseWriter, r *http.Request){
	sess := globalSession1.SessionStart(w, r)

	// w.Write("Welcome " + sess.Get("username"))
	// var wel = sess.Get("username");
	fmt.Println(sess)
	fmt.Println("welcome", sess.Get("username"))
	var wel = fmt.Sprintf("Welcome %q", sess.Get("username"))
	io.WriteString(w,wel)
	
}
func init(){
	globalSession, err := session.NewManager("memory", "gosessionid", 3600)
	// go globalSession.GC()
	globalSession1 = globalSession; 
	if err != nil {
		fmt.Println(err)
		panic("global session creation failed.")
	}
	fmt.Println(globalSession)
}
type User struct{
	username string 
}
func login(w http.ResponseWriter, r *http.Request){
	sess := globalSession1.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		u :=sess.Get("username")
		if u != nil {
			v := fmt.Sprintf("%v", u)
			fmt.Println(v)
			t.Execute(w, &User{username:v} )
		} else {
			// fmt.Println(&User{username:v})
			fmt.Println(u)
			t.Execute(w, sess.Get("username"))
		}
		// fmt.Println(&User{username:v})
		
	} else {
		sess.Set("username", r.Form.Get("username"))
		fmt.Println(sess)
		fmt.Println(sess.Get("username"))

		http.Redirect(w, r, "/", 302)
	}
}
func count(w http.ResponseWriter, r *http.Request){
	sess := globalSession1.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", ct.(int) + 1)
	}
	t, _:= template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}