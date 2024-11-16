package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Page struct {
	Router   *chi.Mux
	tmpl     *template.Template
	tmplName string
	Title    string
}

func main() {

	templates := template.Must(template.ParseGlob("tmpl/*.html"))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	page := &Page{
		tmpl:     templates,
		tmplName: "index.html",
		Title:    "Go Web Template",
	}

	/* for each new page on your website, make one of these!
	basically its the name of the page in your code := &Page{}

	next you'll add the important info
	this one should always be there >> tmpl: templates
	next the name of the html file >> tmplName: "myfile.html"
	Finally, add the title for the page >> Title: "My webpage!"

	you'll need to have the <title> tag in you html set like this
	for the title to work right ||
							    \/
					<title>{{.Title}}</title>

				Look at the html files for refrence
	*/

	testPage := &Page{
		tmpl:     templates,
		tmplName: "test.html",
		Title:    "Test Page!",
	}

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	/*
		You'll also want to add a request handler like the ones below
		example:
		r.Get(("/website_url"), NameOfPageInUrCode.mainHandler)
									/\
									||
			This part is whatever you put above for the := &Page thingy
	*/

	r.Get(("/"), page.mainHandler)
	r.Get(("/test"), testPage.mainHandler)

	http.ListenAndServe(":8080", r)

}

/* to run this webserver, you can type go run main.go in a
command prompt/terminal, or set up air (i forgot how it works but
def way better, look it up!)

to visit your site, type localhost:8080 in your web browser!

*/
