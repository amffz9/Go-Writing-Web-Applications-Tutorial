//This file contains me working through the golang
//web applications tutorial, I do not claim ownership of their ideas
//and will not necessarily follow the tutorial exactly
package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
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

func viewHandler(response http.ResponseWriter, request *http.Request) {
	//Get everything in the path after the first part of the path, which will be "/view/"
	//and put it into title
	title := request.URL.Path[len("/view/"):]

	//We use the title to load an existing page
	//we are ignoring errors for now
	page, _ := loadPage(title)

	//Then we put the page information into an html string then write in out to our response
	fmt.Fprintf(response, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
