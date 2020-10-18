package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["i"]
	if ok && len(keys[0]) >= 1 {
		fmt.Printf("%s | i = %s \n", time.Now().Format("2006-01-02 15:04:05"), keys[0])
	}
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
