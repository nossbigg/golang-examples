package examples

import (
	"fmt"
	"net/http"
)

// StartHelloWorldServer ... starts hello world server
func StartHelloWorldServer() {
	http.HandleFunc("/hello-world", HelloWorldHandler)
	fmt.Print("server started on port 8000")
	http.ListenAndServe(":8000", nil)
}

// HelloWorldHandler ... exported for testing
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(400)
		fmt.Print("request rejected")
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Hello world!"))
}
