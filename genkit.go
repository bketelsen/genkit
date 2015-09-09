package main

import (
	"log"
	"os"
)

var files = []string{"instrumenting", "logging", "registration", "service", "transport"}

func processFile(inputPath string) {
	log.Printf("Processing file %s", inputPath)

	packageName, types := loadFile(inputPath)

	log.Printf("Found genkit types to generate: %#v", types)
	for _, suffix := range files {
		outputPath, err := getRenderedPath(suffix, inputPath)
		if err != nil {
			log.Fatalf("Could not get output path: %s", err)
		}

		output, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("Could not open output file: %s", err)
		}
		defer output.Close()

		if err := render(suffix, output, packageName, types); err != nil {
			log.Fatalf("Could not generate go code: %s", err)
		}
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("genkit: ")

	for _, path := range os.Args[1:] {
		processFile(path)
	}
}
