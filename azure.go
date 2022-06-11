package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func getTokens(flags Flags) []Token {
	params := []string{
		"displayFilterOption=1",
		"createdByOption=3",
		"sortByOption=3",
		"startRowNumber=1",
		"pageSize=100",
		"api-version=7.0-preview.1",
	}

	encodedToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:none", flags.Token)))

	url := fmt.Sprintf(
		"https://vssps.dev.azure.com/%s/_apis/Token/SessionTokens?%s",
		flags.Organization,
		strings.Join(params, "&"))

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	headers := http.Header{
		"Authorization": []string{fmt.Sprintf("Basic %s", encodedToken)},
	}
	req.Header = headers

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tokenList RootObject

	if err := json.Unmarshal(responseData, &tokenList); err != nil {
		fmt.Println("failed to unmarshal:", err)
	}

	var filteredTokens []Token

	for _, token := range tokenList.Value {
		if flags.ValidOnly && !token.IsValid {
			continue
		}

		dateString, err := time.Parse("2006-01-02T15:04:05", strings.Split(token.Expiration, ".")[0])

		if err != nil {
			log.Fatal(err)
		}

		tokenFormatted := Token{
			DisplayName:    token.DisplayName,
			IsValid:        token.IsValid,
			Expiration:     dateString.Format("2006-01-02 15:04:05"),
			Scope:          token.Scope,
			TargetAccounts: token.TargetAccounts,
		}

		if flags.Filter == "" {
			filteredTokens = append(filteredTokens, tokenFormatted)
		} else if token.DisplayName == flags.Filter {
			filteredTokens = append(filteredTokens, tokenFormatted)
		}
	}

	return filteredTokens
}

type RootObject struct {
	Count int     `json:"count"`
	Value []Token `json:"value"`
}

type Token struct {
	DisplayName    string   `json:"displayName"`
	Expiration     string   `json:"validTo"`
	IsValid        bool     `json:"isValid"`
	Scope          string   `json:"scope"`
	TargetAccounts []string `json:"targetAccounts"`
}
