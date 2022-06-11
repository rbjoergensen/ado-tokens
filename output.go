package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func printOutput(tokens []Token, format string) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	if format == "table" {
		fmt.Fprintln(w, "DisplayName\t IsValid\t Expiration\t Scope\t TargetAccounts")
	}

	for _, token := range tokens {
		if !contains([]string{"json", "table"}, format) {
			log.Fatal(format, " is not an option for output")
		}

		if format == "table" {
			fmt.Fprintln(w,
				token.DisplayName, "\t",
				token.IsValid, "\t",
				token.Expiration, "\t",
				token.Scope, "\t",
				token.TargetAccounts)
		}
	}

	w.Flush()

	if format == "json" {
		jsonString, err := json.Marshal(tokens)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(jsonString))
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
