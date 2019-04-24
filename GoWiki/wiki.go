//https://golang.org/doc/articles/wiki/

package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func main() {

	// first code:
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	// second code:
	// http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)

	// third code:
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// func saveHandler(w http.ResponseWriter, r *http.Request) {
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/save/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	//p.save()
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/view/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }
	//p, _ := loadPage(title)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	// 1: can be replaced by template
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)

	renderTemplate(w, "view", p)
}

// func editHandler(w http.ResponseWriter, r *http.Request) {
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/edit/"):]

	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	// 1: can be replaced by template
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<textarea name=\"body\">%s</textarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)

	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// 1.
	// t, _ := template.ParseFiles(tmpl + ".html")
	// t.Execute(w, p)

	// 2.
	// t, err := template.ParseFiles(tmpl + ".html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// err = t.Execute(w, p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])

	}
}
