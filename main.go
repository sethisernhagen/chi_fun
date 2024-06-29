package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	if err := userStoreApi.SetupRoutes(router); err != nil {
		panic(err)
	}
	_ = http.ListenAndServe(":8080", router)

}
