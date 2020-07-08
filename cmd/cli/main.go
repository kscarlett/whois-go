package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kscarlett/whois-go/pkg/whois"
)

func main() {
	err, query, server := handleArgs()
	if err != nil {
		printUsage()
		os.Exit(1)
	}

	err, result := whois.Query(query, server)
	if err == whois.ErrConnectionIssue {
		fmt.Fprintf(os.Stderr, "An issue occurred: %w\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", result)
		os.Exit(0)
	}
}

func handleArgs() (err error, query, server string) {
	switch len(os.Args) {
	case 2:
		return nil, os.Args[1], ""
	case 3:
		return nil, os.Args[1], os.Args[2]
	default:
		return errors.New("incorrect amount of arguments"), "", ""
	}
}

func printUsage() {
	// TODO: add usage message
	// binary <query> [server]
}
