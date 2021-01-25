// Package grep implements a function like UNIX grep - just without the regex.
package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Search for a pattern in files, given certain flags
// - `-n` Print the line numbers of each matching line.
// - `-l` Print only the names of files that contain at least one matching line.
// - `-i` Match line using a case-insensitive comparison.
// - `-v` Invert the program -- collect all lines that fail to match the pattern.
// - `-x` Only match entire lines, instead of lines that contain a match.
func Search(pattern string, flags, files []string) []string {
	// map flags slice to a map so it is easier to check if a flag is active
	config := make(map[string]bool)
	for _, f := range flags {
		config[f] = true
	}

	result := []string{}
	for _, fname := range files {
		file, err := os.Open(fname)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		lineNumber := 0
		// set up a new SplitFunc for bufio.Scanner that updates the lineNumber:
		scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			advance, token, err = bufio.ScanLines(data, atEOF)
			lineNumber++
			return
		}
		scanner.Split(scanLines)

		for scanner.Scan() {
			currentLine := scanner.Text()
			normalizedPattern := pattern
			prefix := ""

			if len(files) > 1 {
				if strSliceContains(result, fname) {
					continue // only want the file name
				}
				prefix = fmt.Sprintf("%s:", file.Name())
			}

			if config["-n"] { // add line number?
				prefix = fmt.Sprintf("%s%d:", prefix, lineNumber)
			}

			if config["-i"] { // case-insensitive match?
				currentLine = strings.ToUpper(currentLine)
				normalizedPattern = strings.ToUpper(normalizedPattern)
			}

			if !config["-x"] && !config["-v"] { // neither want entire line match nor inverted behaviour?
				if strings.Contains(currentLine, normalizedPattern) {
					updateResult(&result, config["-l"], fname, prefix, scanner.Text())
				}
			}

			if config["-x"] && !config["-v"] { // only entire lines but not inverted?
				if currentLine == normalizedPattern {
					updateResult(&result, config["-l"], fname, prefix, scanner.Text())
				}
			}

			if config["-v"] { // inverted behaviour?
				if config["-x"] { // only entire line match?
					if currentLine != normalizedPattern {
						updateResult(&result, config["-l"], fname, prefix, scanner.Text())
					}
				} else { // part of line match?
					if !strings.Contains(currentLine, normalizedPattern) {
						updateResult(&result, config["-l"], fname, prefix, scanner.Text())
					}
				}
			}

		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return result
}

// strSliceContains - a helper to check if a slice of strings contains a certain string
func strSliceContains(list []string, what string) bool {
	for _, entry := range list {
		if entry == what {
			return true
		}
	}
	return false
}

// updateResult - a helper to extend the results slice in-place (result passed as pointer)
func updateResult(result *[]string, printFname bool, fname string, prefix string, line string) {
	if printFname {
		*result = append(*result, fmt.Sprintf("%s", fname))
	} else {
		*result = append(*result, fmt.Sprintf("%s%s", prefix, line))
	}
}
