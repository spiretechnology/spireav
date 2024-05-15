package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

const (
	// FilterDir is the directory where the filtergen tool will generate the filter packages.
	FilterDir = "filter"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Pull the list of filters from ffmpeg
	filters, err := listFilters(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Generate the filter packages
	if err := generateFilters(filters); err != nil {
		log.Fatal("error generating filters: ", err)
	}
}
