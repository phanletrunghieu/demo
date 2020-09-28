package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	serviceHttp "github.com/phanletrunghieu/demo/gqlgen/http"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load .env by errors: %v", err))
	}

	// setup port
	defaultPort := "3000"
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler()
	}

	log.Printf("Server is running on http://localhost:%s", port)

	err = http.ListenAndServe(":"+port, h)
	if err != nil {
		log.Fatal(err)
	}
}
