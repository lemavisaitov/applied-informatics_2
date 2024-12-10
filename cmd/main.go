package main

import (
	"fmt"
	"github.com/lemavisaitov/applied-informatics_2/internal/fetcher"
	"github.com/lemavisaitov/applied-informatics_2/internal/processor"
	"github.com/lemavisaitov/applied-informatics_2/internal/validator"
)

func main() {
	fmt.Println("Choose an option:")
	fmt.Println("1. Validate user input")
	fmt.Println("2. Search for IPv4 addresses on a webpage")
	fmt.Println("3. Search for IPv4 addresses in a file")
	fmt.Print("Enter your choice: ")

	var choice int
	_, _ = fmt.Scanln(&choice)

	ipv4Validator := validator.NewIPv4Validator()
	switch choice {
	case 1:
		fmt.Print("Enter a string to validate: ")
		var input string
		_, _ = fmt.Scanln(&input)
		if ipv4Validator.Validate(input) {
			fmt.Println("Valid IPv4 address!")
		} else {
			fmt.Println("Invalid IPv4 address.")
		}
	case 2:
		fmt.Print("Enter the URL: ")
		var url string
		_, _ = fmt.Scanln(&url)

		proc := processor.NewProcessor(&fetcher.URLFetcher{}, ipv4Validator)
		matches, err := proc.Process(url)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Found IPv4 addresses:", matches)
		}
	case 3:
		fmt.Print("Enter the file path: ")
		var filePath string
		_, _ = fmt.Scanln(&filePath)

		proc := processor.NewProcessor(&fetcher.FileFetcher{}, ipv4Validator)
		matches, err := proc.Process(filePath)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Found IPv4 addresses:", matches)
		}
	default:
		fmt.Println("Invalid choice.")
	}
}
