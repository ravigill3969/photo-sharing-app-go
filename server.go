package main

import (
	"fmt"
	"net/http"

	"github.com/ravigill3969/backend/middlewares"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/water", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("water"))
	})

	port := ":8080"
	fmt.Println("server is running on port", port)

	middlewareHandlers := middlewares.Cors(mux)

	err := http.ListenAndServe(port, middlewareHandlers)

	if err != nil {
		fmt.Println("error stating server", err)
		return
	}

}
