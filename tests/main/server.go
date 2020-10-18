package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["i"]

	if ok && len(keys[0]) >= 1 {
		fmt.Println("i = " + string(keys[0]))
	}
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
