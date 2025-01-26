package main

import (
	"fmt"
	_ "net/http"
	"time"

	_ "github.com/go-chi/chi/middleware"
	_ "github.com/go-chi/chi/v5"
	"github.com/ikotun/llmxp/internals/handlers"
	"github.com/joho/godotenv"
)

func main() {

	//env checks
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return

	}

	now := time.Now()
	fmt.Println(time.Unix(now.Unix(), 0))

	handlers.InitOpenAI()
	fmt.Println(time.Unix(now.Unix(), 0))

	// r := chi.NewRouter()
	//
	// r.Use(middleware.Logger)
	//
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	// r.Post("/process-message", handlers.ProcessMessage)
	//
	// fmt.Printf("Server is running on port 6600\n")
	//
	// err = http.ListenAndServe(":6600", r)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
