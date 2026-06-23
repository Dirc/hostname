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

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
}

func hostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s\n", hostname)
	_, err = fmt.Fprintf(w, "Hostname: %s", hostname)
	if err != nil {
		fmt.Printf("Failed to write response: %v\n", err)
	}
}

func env(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("FOO")
	var err error
	if env == "" {
		_, err = fmt.Fprintf(w, "Environment variable FOO is not found")
	} else {
		_, err = fmt.Fprintf(w, "FOO=%s", env)
	}
	if err != nil {
		fmt.Printf("Failed to write response: %v\n", err)
	}
}
