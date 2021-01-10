package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tkircsi/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application")
	file := flag.String("file", "story1.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(*story)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
