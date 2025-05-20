package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/songwaad/learning-go/jokevis"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)
	jokevis.SayJokevis()
	jokevis.SayTest()
}
