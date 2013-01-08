package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "statics/index.html")
}

func main() {
	ip := os.Getenv("OPENSHIFT_INTERNAL_IP")
	http.HandleFunc("/", index)
	http.Handle("/docs", http.FileServer(http.Dir("statics")))
	http.ListenAndServe(fmt.Sprintf("%s:8080", ip), nil)
}
