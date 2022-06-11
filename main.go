package main

func main() {
	flags := loadFlags()

	tokens := getTokens(flags)

	printOutput(tokens, flags)
}
