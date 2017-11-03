package main

import (
	"io/ioutil"
	"fmt"
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

func main() {
	//Make a page to test the saving functionality
	page1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	page1.save()

	//Try to load said page from disk
	page2, _  := loadPage("TestPage")

	//Then print that page to the console
	fmt.Println(string(page2.Body))
}
