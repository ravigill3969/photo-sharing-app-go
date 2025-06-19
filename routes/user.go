package routes

import (
	"net/http"

	"github.com/ravigill3969/backend/controller"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", controller.Register)
}
