// filter_code filters code lines from a file by looking for punctuation.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// htmlRegExp matches lines that start and end with HTML.
var htmlRegExp = regexp.MustCompile(`^</?\\w.*>$`)

// markRegExp matches punctuation marks.
var markRegExp = regexp.MustCompile(`\pP`)

// specialRegExp matches special code marks.
var specialRegExp = regexp.MustCompile(`[=;{}]`)

// wordRegExp matches words
var wordRegExp = regexp.MustCompile(`\w+`)

// camelRegExp matches camelcased words
var camelRegExp = regexp.MustCompile(`[a-z][A-Z]`)

func main() {
	// Parse flags.
	mode := flag.String("type", "dupes", "Type of filter to use")
	flag.Parse()

	// Check argument.
	if len(flag.Args()) < 1 {
		log.Fatal("Missing file argument")
	}
	filename := flag.Arg(0)

	// Choose program function.
	if *mode == "code" {
		filterCode(filename)
	} else if *mode == "dupes" {
		filterDupes(filename)
	} else {
		log.Fatal("Unrecognized type")
	}
}

// filterCode filters out code lines.
func filterCode(filename string) {
	// Open file.
	hdl, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Scan file line by line, removing code.
	scanner := bufio.NewScanner(hdl)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if !isCode(line) {
			fmt.Println(line)
		}
	}
}

// filterDupes filters out duplicate lines and sequential blank lines.
func filterDupes(filename string) {
	// Open file.
	hdl, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Scan file line by line, removing duplicates.
	snippets := make(map[string]bool)
	scanner := bufio.NewScanner(hdl)

	// Keep track of whether the last line was empty or not.
	wasEmpty := false
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// Trim white space on right.
		line = strings.TrimRight(line, " \t")

		// Use lowercase snippet to ignore case.
		snippet := strings.ToLower(line)

		// Remove spaces and tabs to ignore them.
		snippet = strings.Replace(snippet, " ", "", -1)
		snippet = strings.Replace(snippet, "\t", "", -1)

		// Get a snippet of up to 255 characters.
		n := len(snippet)
		if n > 255 {
			n = 255
		}

		// Print the line if it was empty or not recognized.
		if n == 0 {
			if !wasEmpty {
				fmt.Println()
			}
			wasEmpty = true
		} else if !snippets[snippet] {
			fmt.Println(line)
			snippets[snippet] = true
			wasEmpty = false
		}
	}
}

// isCode checks if a line is code.
func isCode(line string) bool {
	// Check for a line that starts with and ends with HTML or XML.
	if htmlRegExp.MatchString(line) {
		return true
	}

	// See if the counts of punctuation and code words outweight normal words.
	markCnt := len(markRegExp.FindAllString(line, -1))
	specialCnt := len(specialRegExp.FindAllString(line, -1))
	wordCnt := len(wordRegExp.FindAllString(line, -1))
	camelCnt := len(camelRegExp.FindAllString(line, -1))
	return markCnt*2+specialCnt*3 > wordCnt-camelCnt
}
