package main

import (
	"fmt"
	"os"
	"io"
	"flag"

	"git.dominic-ricottone.com/~dricottone/parcels/common"
)

func find_in_stream(reader io.Reader, target int) {
	url, err := common.PullFromReader(reader, target, 0)
	if err != nil {
		fmt.Printf("internal error - %v\n", err)
		os.Exit(1)
	}

	fmt.Println(url)
}

func parse_stream(reader io.Reader) {
	content, urls, err := common.ParseFromReader(reader, 0)
	if err != nil {
		fmt.Printf("internal error - %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s", content)
	fmt.Printf("%s", urls)
}

func parse_file(filename string) {
	// Check file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("cannot read file '%s'\n", filename)
		os.Exit(1)
	}
	defer file.Close()

	parse_stream(file)
}

func main() {
	// Check STDIN
	_, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("cannot read input")
		os.Exit(1)
	}

	// Look for arguments
	var index = flag.Int("n", -1, "print Nth URL")
	flag.Parse()

	// Parse
	if *index < 0 {
		parse_stream(os.Stdin)
	} else {
		find_in_stream(os.Stdin, *index)
	}
}

