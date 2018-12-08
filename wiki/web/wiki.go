package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"fmt"
	//"errors"
)

type Page struct{
	Title string
	Body []byte
	NewBody template.HTML
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html","FrontPage.html"))
//^(start string)...(end string)$
var vaildPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var pageNameRegexp = regexp.MustCompile("\\[([a-zA-Z0-9]+)\\]")
var pagePath = "data/"

func (p *Page) save() error {
	filename := pagePath + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := pagePath + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	
	p, err := loadPage(title)
	
	if title == "FrontPage" {
		renderTemplate(w, "FrontPage", p)
		return 
	}
	
	if err != nil {
		http.Redirect(w, r, "/edit/" +  title, http.StatusFound)
		return
	}
	/*
		[PageName] -> <a href="/view/PageName">PageName</a>
	*/
	escapeBody := []byte(template.HTMLEscapeString(string(p.Body)))
	p.NewBody = template.HTML(pageNameRegexp.ReplaceAllFunc(escapeBody, func(s []byte) []byte {
			m := string(s[1 : len(s)-1])
			return []byte("<a href=\"/view/"+ m + "\">" + m +"</a>")
	}))
	
	renderTemplate(w, "view", p)//
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/edit/"):]
	
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}	
	//t, _ := template.ParseFiles("edit.html")
	//t.Execute(w, p)
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/save/"):]
	
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
    	fmt.Println("error occurred")
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
/*
	change closuer
	
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}
*/

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := vaildPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			//http.NotFound(w, r)
			http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
			return 
		}
		fn(w, r, m[2])
	}
}


func main() {
	//TODO: add /view/frontPage/ and linking page 
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/search/", searchHandler)
	http.ListenAndServe(":8080", nil)
}



