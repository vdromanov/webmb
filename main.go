package main

import (
	"fmt"
	"net/http"

	views "webmb/views"
)

func process(w http.ResponseWriter, r *http.Request) { //Add parsing here
	r.ParseForm()
	//	fmt.Fprintln(w, r.PostForm)
	fmt.Println(r.FormValue("portName"))
	fmt.Println(r.PostForm)
	http.Redirect(w, r, "/index", http.StatusSeeOther)

}

func main() {
	server := http.Server{
		Addr: ":4040",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/index", views.IndexView)
	http.HandleFunc("/", views.ConfigurePortView)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	server.ListenAndServe()
}
