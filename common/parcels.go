package common

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"regexp"
)

var urlPattern = regexp.MustCompile(UrlPattern)

// Modify a string such that replacement occupies it from the beginning index
// to the end index.
func replace(str string, beginning int, end int, replacement int) string {
	return str[:beginning] + fmt.Sprintf("[%d]", replacement) + str[end:]
}

// Pull a URL from a scanner.
func pullFromScanner(scanner *bufio.Scanner, target int) (string, error) {
	target_url := ""
	count_urls_skipped := 0

	for scanner.Scan() {
		// find all matches (and count of matches) on this line
		line := scanner.Text()
		matches := urlPattern.FindAllStringIndex(line, -1)
		count_urls_after_line := count_urls_skipped + len(matches)

		// if target url is on this line, pull it from matches
		if target < count_urls_after_line {
			target_beg := matches[target - count_urls_skipped][0]
			target_end := matches[target - count_urls_skipped][1]
			target_url = line[target_beg:target_end]
			break
		}

		// else update count skipped and go to next line
		count_urls_skipped = count_urls_after_line
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return target_url, err
	}

	return target_url, nil
}

// Parse URLs from a scanner. Return two strings: the modified content of the
// scanner, and a list of URLs. Use offset to indicate how many URLs preceded
// this scanner.
func parseFromScanner(scanner *bufio.Scanner, offset int) (string, string, error) {
	var content strings.Builder
	var urls strings.Builder
	cursor := offset

	for scanner.Scan() {
		// find all matches (and count of matches) on this line
		line := scanner.Text()
		matches := urlPattern.FindAllStringIndex(line, -1)
		count := len(matches)
		var new_urls = make([]string, count)

		// looping backwards, extract each URL and replace it in the
		// content
		for i := count - 1; i >= 0; i-- {
			target_beg := matches[i][0]
			target_end := matches[i][1]
			new_urls[i] = line[target_beg:target_end]
			line = replace(line, target_beg, target_end, cursor+i)
		}

		// update the content
		content.WriteString(line)
		content.WriteString("\n")

		// update the list of urls
		for i, url := range new_urls {
			urls.WriteString(fmt.Sprintf("[%d] %s\n", cursor+i, url))
		}

		// update the cursor
		cursor += count
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return content.String(), urls.String(), err
	}

	return content.String(), urls.String(), nil
}

// Pull a URL from a reader. Use offset to indicate how many URLs preceded this
// reader.
func PullFromReader(reader io.Reader, target int, offset int) (string, error) {
	return pullFromScanner(bufio.NewScanner(reader), target-offset)
}

// Pull a URL from a string. Use offset to indicate how many URLs preceded this
// string.
func PullFromString(str string, target int, offset int) (string, error) {
	return PullFromReader(strings.NewReader(str), target, offset)
}

// Parse URLs from a reader. Return two strings: the modified content of the
// reader and a list of URLs. Use offset to indicate how many URLs preceded
// this reader.
func ParseFromReader(reader io.Reader, offset int) (string, string, error) {
	return parseFromScanner(bufio.NewScanner(reader), offset)
}

// Parse URLs from a string. Return two strings: the modified content of the
// original string and a list of URLs. Use offset to indicate how many URLs
// preceded the original string.
func ParseFromString(str string, offset int) (string, string, error) {
	return ParseFromReader(strings.NewReader(str), offset)
}

