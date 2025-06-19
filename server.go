package main

import (
	"fmt"
	"net/http"

	"github.com/ravigill3969/backend/database"
	"github.com/ravigill3969/backend/middlewares"
	"github.com/ravigill3969/backend/routes"
)

func main() {

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}
	defer db.Close()

	port := ":8080"
	fmt.Println("server is running on port", port)

	middlewareHandlers := middlewares.Cors(mux)

	err = http.ListenAndServe(port, middlewareHandlers)

	if err != nil {
		fmt.Println("error stating server", err)
		return
	}

}
