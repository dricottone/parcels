package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"regexp"
	"flag"
)

func Replacement(i int) string {
	return fmt.Sprintf("[%d]", i)
}

func find_in_stream(reader io.Reader, target int) {
	// Create scanner from reader
	input := bufio.NewScanner(reader)

	// Initialize state
	re := regexp.MustCompile(UrlPattern)
	count := 0

	// Parse and print
	for input.Scan() {
		line := input.Text()
		line_indices := re.FindAllStringIndex(line, -1)
		count_after := count + len(line_indices)

		if target < count_after {
			beg, end := line_indices[target-count][0], line_indices[target-count][1]
			fmt.Println(line[beg:end])
			break
		}

		count = count_after
	}

	// Check for scanner errors
	if err := input.Err(); err != nil {
		fmt.Printf("internal error - %v\n", err)
		os.Exit(1)
	}
}

func parse_stream(reader io.Reader) {
	// Create scanner from reader
	input := bufio.NewScanner(reader)

	// Initialize state
	re := regexp.MustCompile(UrlPattern)
	var parcels [999]string //assuming that never need 1000+ URLs
	offset := 0

	// Parse, modify, and print
	for input.Scan() {
		line := input.Text()

		line_indices := re.FindAllStringIndex(line, -1)
		for i := len(line_indices)-1; i >= 0; i-- {
			beg, end := line_indices[i][0], line_indices[i][1]
			parcels[offset+i] = line[beg:end]
			line = line[:beg] + Replacement(offset+i) + line[end:]
		}

		fmt.Println(line)

		offset = offset + len(line_indices)
	}

	// Print postscript
	fmt.Printf("\nURLs:\n")
	for index, url := range parcels {
		if url == "" {
			break
		}
		fmt.Printf("[%d] %s\n", index, url)
	}

	// Check for scanner errors
	if err := input.Err(); err != nil {
		fmt.Printf("internal error - %v\n", err)
		os.Exit(1)
	}
}

func parse_file(filename string) {
	// Check file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("cannot read file '%s'\n", filename)
		os.Exit(1)
	}
	defer file.Close()

	// Parse
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

