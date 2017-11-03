//This file contains me working through the golang
//web applications tutorial, I do not claim ownership of their ideas
//and will not necessarily follow the tutorial exactly
package main

import (
	"io/ioutil"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
func saveHandler(response http.ResponseWriter, request *http.Request) {

}

//edit handler loads the page(or an empty one if it doesn't exist) and displays an html form
func editHandler(response http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(response, "edit", p)
}

func viewHandler(response http.ResponseWriter, request *http.Request) {
	//Get everything in the path after the first part of the path, which will be "/view/"
	//and put it into title
	title := request.URL.Path[len("/view/"):]

	//We use the title to load an existing page
	//we are ignoring errors for now
	page, err := loadPage(title)

	if err != nil {
		page = &Page{Body: []byte("Empty"), Title: "This view does not exist"}
	}

	renderTemplate(response, "view", page)
}

//a function to write a template for the server
func renderTemplate(response http.ResponseWriter, templateString string, page *Page) {
	t, _ := template.ParseFiles(templateString + ".html")
	t.Execute(response, page)
}
//save operates on a page type and saves it to disk, using the page title as the filename
//and saving the body of the page as the file content
func (p *Page) save() error {
	filename := p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)
}

//loadPage takes the title of a page as a parameter and then loads the information that has been saved into
//a Page structure
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	//what to do if the file has failed to load?
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
