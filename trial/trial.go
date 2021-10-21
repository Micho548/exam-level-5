package trial

import (
	"fmt"
	"log"
	"net/http"
)

// handler
func Hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Method used:", r.Method)
	fmt.Fprintln(rw, "hello world")
}

func SubMain() {

	// Router
	http.HandleFunc("/", Hello)
	fmt.Println("server running on port 8000")

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
