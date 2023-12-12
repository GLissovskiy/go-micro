package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const webPort = "8080"

func main() {
	app := Config{}

	log.Println("Starting mail service on port: ", webPort)

	srv := &http.Server{

		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}