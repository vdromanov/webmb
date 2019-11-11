package views

import (
	"net/http"
	"text/template"
)

//Caching templates at init
var templates = make(map[string]*template.Template)

var configurePortFormNames = map[string]string{
	"name":   "portName",
	"bd":     "baudrate",
	"parity": "parity",
	"bits":   "dataBits",
	"stop":   "stopbits",
}

//General representation of all rendered HTML templates
type pageToRender struct {
	Title          string
	FormFieldNames *map[string]string
	ActionRoute    string //Where sent submit
}

func init() {
	for _, templ := range []string{"templates/index.html", "templates/setup_port.html"} {
		t, err := template.ParseFiles(templ)
		if err != nil {
			panic(err)
		}
		templates[templ] = t
	}
}

func RenderTemplate(w http.ResponseWriter, templateFname string, p *pageToRender) {
	w.Header().Set("Content-type", "text/html")
	if err := templates[templateFname].Execute(w, p); err != nil {
		http.Error(w, "Incorrect template: "+templateFname, http.StatusInternalServerError)
	}
}

func ConfigurePortView(w http.ResponseWriter, r *http.Request) {
	p := pageToRender{
		Title:          "Port Settings",
		FormFieldNames: &configurePortFormNames,
		//		ActionRoute:    "http://127.0.0.1:4040/process",
		ActionRoute: "/process",
	}
	RenderTemplate(w, "templates/setup_port.html", &p)
}

func IndexView(w http.ResponseWriter, r *http.Request) {
	p := pageToRender{
		Title:       "Testing",
		ActionRoute: "/process",
		//		ActionRoute: "http://127.0.0.1:4040/process",
	}
	RenderTemplate(w, "templates/index.html", &p)
}
