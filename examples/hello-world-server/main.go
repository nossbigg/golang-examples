package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", helloWorldHandler)
	fmt.Print("server started on port 8000")
	http.ListenAndServe(":8000", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Print("request rejected")
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Hello world!"))
}
