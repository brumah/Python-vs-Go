package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/brunerMatthew/Python-vs-Go/queries"
)

func main() {
	http.HandleFunc("/getquery", queryHandler)
	http.HandleFunc("/", rootHandler)

	port := ":8080"
	fmt.Printf("Server is listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	initialText := "Hello"
	tmpl.Execute(w, initialText)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	query := r.URL.Query()
	id := query.Get("id")
	var text string

	switch id {
	case "1":
		text = queries.BlueButton()
		fmt.Println(time.Since(start))
	case "2":
		text = queries.GreenButton()
		fmt.Println(time.Since(start))
	case "3":
		text = queries.RedButton()
		fmt.Println(time.Since(start))
	case "4":
		text = queries.PurpleButton()
		fmt.Println(time.Since(start))
	case "5":
		text = queries.ClearButton()
		fmt.Println(time.Since(start))
	}

	tmpl, _ := template.New("t").Parse(text)
	tmpl.Execute(w, nil)
}
