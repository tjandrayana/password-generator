package backend

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Test() {

	fmt.Println("Listening ... ")
	fs := http.FileServer(http.Dir("./front/"))
	http.Handle("/front/", http.StripPrefix("/front/", fs))

	http.HandleFunc("/generator/", viewHandler)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	p := LoadPageSrizon("TestPage")
	fmt.Fprintf(w, string(p.Body))
}

type Page struct {
	Title string
	Body  []byte
}

func LoadPageSrizon(title string) *Page {
	// filename := "./front/html/guitar.html"

	filename := "./front/index.html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return &Page{Title: title, Body: body}
}
