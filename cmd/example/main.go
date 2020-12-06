package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	lab2 "github.com/teramont/go-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File to output the result")
)

func main() {
	flag.Parse()
	var reader io.Reader
	var writer io.Writer

	if len(*inputExpression) > 0 {
		reader = strings.NewReader(*inputExpression)
	} else if len(*inputFile) > 0 {
		file, err := os.Open(*inputFile)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		reader = file
	} else {
		log.Fatal("Reader interface not defined")
	}

	if len(*outputFile) > 0 {
		file, err := os.Create(*outputFile)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		writer = file
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()
	if err != nil {
		log.Fatal(err)
	}
}
