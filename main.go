package main

import (
	"net/http"
	"log"
)

func main() {
	// Setting the directory for static files in the "static" folder
	fs := http.FileServer(http.Dir("./static"))

	// Serving the static directory on the route "/static/"
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setting the home route
	http.HandleFunc("/", homeHandler)

	// Starting the server on port 8080 and catching any errors
	log.Println("Serving on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// homeHandler serves the index.html file
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}