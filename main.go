package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(msg string) string {
	return fmt.Sprintf("<b>%s</b>", msg)
}

func main() {
	msg := greeting("Code.education Rocks!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<body>%s</body>", msg)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
