package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	//env checks
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return

	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	fmt.Printf("Server is running on port 6600\n")

	err = http.ListenAndServe(":6600", r)
	if err != nil {
		fmt.Println(err)
	}
}
