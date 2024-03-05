package main

import (
	"github.com/wojcikt/trucktel-go/internal/server"
	"log"
)

func main() {
	srv := server.New()
	log.Fatalln(srv.ListenAndServe())
}
