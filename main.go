package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hostname)
	router.HandleFunc("/env", env)
	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}

func hostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprintf(w, "Hostname: %s", hostname )
}

func env(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("FOO")
	if env == "" {
		fmt.Fprintf(w, "Environment variable FOO is not found")
	} else {
    fmt.Fprintf(w, "FOO=%s", env )
	}
}