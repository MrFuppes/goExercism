// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey returns Bob's answer
func Hey(remark string) string {
	// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
	var response string

	remark = strings.TrimSpace(remark)

	isQuest := strings.HasSuffix(remark, "?")

	// the following logic will fail for characters like á, ö etc.
	reg, _ := regexp.Compile("[a-zA-Z]")
	hasChars := len(reg.FindAllStringIndex(remark, -1)) > 0
	isShout := hasChars && strings.Compare(remark, strings.ToUpper(remark)) == 0

	switch {
	// Bob answers 'Sure.' if you ask him a question, such as "How are you?".
	case isQuest && !isShout:
		response = "Sure."
	// He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
	case !isQuest && isShout:
		response = "Whoa, chill out!"
	// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
	case isQuest && isShout:
		response = "Calm down, I know what I'm doing!"
	// He says 'Fine. Be that way!' if you address him without actually saying
	// anything.
	case len(remark) == 0:
		response = "Fine. Be that way!"
	// He answers 'Whatever.' to anything else.
	default:
		response = "Whatever."
	}

	return response
}
