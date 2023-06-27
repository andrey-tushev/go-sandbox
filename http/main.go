package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func access(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !r.URL.Query().Has("ok") {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Access denied")
			return
		}

		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func main() {
	http.HandleFunc("/foo", access(logging(foo)))
	http.HandleFunc("/bar", access(logging(bar)))
	http.ListenAndServe(":8087", nil)
}
