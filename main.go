package main

func main() {
	flags := loadFlags()
	tokens := getTokens(flags.Token, flags.Organization, flags.Filter)
	printOutput(tokens, flags.Output)
}
