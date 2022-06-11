package main

import (
	"flag"
	"log"
)

func loadFlags() Flags {
	tokenPtr := flag.String(
		"token",
		"",
		"The token used to access the organization (*Required)")
	organizationPtr := flag.String(
		"org",
		"",
		"The organization the token was created in (*Required)")
	filterPtr := flag.String(
		"name",
		"",
		"The name of the token(s) to find")
	outputPtr := flag.String(
		"output",
		"table",
		"The format in which to output the results(table, json)")

	flag.Parse()

	if *tokenPtr == "" {
		log.Fatal("--token flag is missing")
	}
	if *organizationPtr == "" {
		log.Fatal("--org flag is missing")
	}

	return Flags{
		Token:        *tokenPtr,
		Organization: *organizationPtr,
		Filter:       *filterPtr,
		Output:       *outputPtr,
	}
}

type Flags struct {
	Token        string
	Organization string
	Filter       string
	Output       string
}
