package main

import "net/http"

func (p *Page) mainHandler(w http.ResponseWriter, r *http.Request) {
	p.tmpl.ExecuteTemplate(w, p.tmplName, p)
}

/* nothing too interesting here rn, but you can set up a url to run some
kind of code here! in my other project i had a custom handler that would
run code to grab definitions off of the internet!
*/
