package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	ip := os.Getenv("OPENSHIFT_INTERNAL_IP")
	http.Handle("/", http.FileServer(http.Dir("statics")))
	http.ListenAndServe(fmt.Sprintf("%s:8080", ip), nil)
}
