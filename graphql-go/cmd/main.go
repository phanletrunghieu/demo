package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/phanletrunghieu/demo/graphql-go/config/database/pg"
	serviceHttp "github.com/phanletrunghieu/demo/graphql-go/http"
	"github.com/phanletrunghieu/demo/graphql-go/resolver"
	"github.com/phanletrunghieu/demo/graphql-go/schema"
	"github.com/phanletrunghieu/demo/graphql-go/service"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load .env by errors: %v", err))
	}

	// setup addr
	httpAddr := ":" + os.Getenv("PORT")

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	// connect db
	db, closeDB := pg.New(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	defer closeDB()

	// init db service
	serv, err := service.NewService(db)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	// init schema
	graphqlSchema := schema.String()
	graphqlResolver := resolver.NewResolver(serv)

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(graphqlSchema, graphqlResolver)
	}

	http.ListenAndServe(httpAddr, h)
}
