package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa Go is neat!</h1>")
	fmt.Fprintf(w, "<p>Go is Fast!!</p>")
	fmt.Fprintf(w, "<p>... and simple</p>")
	fmt.Fprintf(w, "<p>You %s even add  %s<p>", "can", "<strong>strong</strong> characters")
	fmt.Fprintf(w, `
	<h6>You can even do</h6>
	<h5>multiple lines</h6>
	<h4>in one %s</h4>`, "formatted print")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Harrison Kinsley")

}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)

}
