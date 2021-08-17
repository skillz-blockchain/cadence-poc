package main

import (
	"cadence-poc/api/server"
	"cadence-poc/cadence/workers"
	"fmt"
	"log"
	"os"
)

func main() {
	host, port := os.Getenv("API_HOST"), os.Getenv("API_PORT")
	if host == "" || port == "" {
		log.Fatalln("API_HOST | API_PORT missing in env")
	}

	var worker workers.CadenceAdapter
	worker.Setup()

	s := server.NewServer(worker)
	log.Println("server created...")
	if err := s.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Fatalln(err)
	}
}
