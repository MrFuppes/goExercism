// Package markdown - the refactoring exercise
// Code is harder to read than to write -> rewriting may be easier than refactoring...
package markdown

import (
	"fmt"
	"strings"
)

// setStyle - bold and italics
func setStyle(s string) string {
	countDunder := strings.Count(s, "__")
	if countDunder > 0 && countDunder%2 == 0 { // there must be two of them...
		for i := 0; i < countDunder; i = i + 2 {
			s = strings.Replace(s, "__", "<strong>", 1)
			s = strings.Replace(s, "__", "</strong>", 1)
		}
	}
	countDunder, count := strings.Count(s, "__"), strings.Count(s, "_")
	if count > 0 && count%2 == 0 && countDunder == 0 { // must not be dunders...
		for i := 0; i < count; i = i + 2 {
			s = strings.Replace(s, "_", "<em>", 1)
			s = strings.Replace(s, "_", "</em>", 1)
		}
	}

	return s
}

// setHeading - set heading html tags
func setHeading(s string) string {
	if !strings.HasPrefix(s, "#") {
		return s
	}
	var count int
	for _, c := range s { // count number of adjacent #
		if c == '#' {
			count++
		} else {
			break
		}
	}

	return fmt.Sprintf("<h%d>%s</h%d>", count, s[count+1:len(s)], count)
}

// setUnorderedList - set tags to specify an unordered list
func setUnorderedList(lines []string, line string, currentIdx int) string {
	if currentIdx == 0 { // first line, need a <ul> to start
		line = "<ul><li>" + string(line[2:len(line)]) + "</li>"
	}
	if currentIdx > 0 {
		if strings.HasPrefix(lines[currentIdx-1], "* ") {
			line = "<li>" + string(line[2:len(line)]) + "</li>"
		} else { // previous line was not part of list, so we also need a <ul> at start
			line = "<ul><li>" + string(line[2:len(line)]) + "</li>"
		}
	}
	if currentIdx < len(lines)-1 {
		if !strings.HasPrefix(lines[currentIdx+1], "* ") {
			line += "</ul>" // next line is no list anymore, need to terminate
		}
	}
	if currentIdx == len(lines)-1 {
		line += "</ul>" // there is no next line, also need to terminate
	}
	return line
}

// Render translates markdown to HTML
func Render(markdown string) string {
	lines := strings.Split(markdown, "\n")
	html := make([]string, len(lines))
	for i, line := range lines {
		// start with formatting..
		line = setStyle(line)
		// now it can be a heading, a list or a normal paragraph
		switch {
		case strings.HasPrefix(line, "#"):
			line = setHeading(line)
		case strings.HasPrefix(line, "* "):
			line = setUnorderedList(lines, line, i)
		default:
			line = "<p>" + line + "</p>"
		}
		html[i] = line
	}

	return strings.Join(html, "")
}

// --> this only happens to work for the given test cases. what if there are multiple bold itmes?
// markdown = strings.Replace(markdown, "__", "<strong>", 1)
// markdown = strings.Replace(markdown, "__", "</strong>", 1)
// markdown = strings.Replace(markdown, "_", "<em>", 1)
// markdown = strings.Replace(markdown, "_", "</em>", 1)
// --> replaced with function setStyle

// header := 0 // --> this is actually a counter, misleading name?
// pos := 0
// list := 0 // --> this is actually a counter, misleading name?
// html := ""
// for {
// 	char := markdown[pos]
// 	if char == '#' {
// 		for char == '#' {
// 			header++
// 			pos++
// 			char = markdown[pos]
// 		}
// 		html += fmt.Sprintf("<h%d>", header)
// 		pos++
// 		continue
// 	}
// 	if char == '*' {
// 		if list == 0 {
// 			html += "<ul>"
// 		}
// 		html += "<li>"
// 		list++
// 		pos += 2
// 		continue
// 	}
// 	if char == '\n' {
// 		if list > 0 {
// 			html += "</li>"
// 		}
// 		if header > 0 {
// 			html += fmt.Sprintf("</h%d>", header)
// 			header = 0
// 		}
// 		pos++
// 		continue
// 	}
// 	html += string(char)
// 	pos++
// 	if pos >= len(markdown) {
// 		break
// 	}
// }

// if header > 0 {
// 	return html + fmt.Sprintf("</h%d>", header)
// }

// if list > 0 {
// 	return html + "</li></ul>"
// }

// return html //"<p>" + html + "</p>"
