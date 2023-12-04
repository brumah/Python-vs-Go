package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/brunerMatthew/Python-vs-Go/queries"
)

var latencyValues []string

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
		text, _ = queries.BlueButton()
	case "2":
		text, _ = queries.RedButton()
	case "3":
		text, _ = queries.PurpleButton()
	case "4":
		text = queries.ClearButton()
	case "5":
		text = queries.ExportButton(latencyValues)
		latencyValues = []string{}
	}

	tmpl, _ := template.New("t").Parse(text)
	tmpl.Execute(w, nil)
	fmt.Println(time.Since(start))
	if id != "6" {
		latencyValues = append(latencyValues, fmt.Sprintf("%.2f", float64(time.Since(start))/1000000))
	}
}
