package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v45/github"
)

func main() {

	ctx := context.Background()
	client := github.NewClient(nil)

	arg := "silvanocostanzo"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	// TODO: check rate limiting before fetching the repositories
	// https://pkg.go.dev/github.com/google/go-github/v45/github#hdr-Rate_Limiting

	reps, _, err := client.Repositories.List(ctx, arg, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range reps {
		fmt.Printf("%s\n", *repo.Name)
	}
}
