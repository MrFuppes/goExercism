package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

// Team - a struct to hold results of a team
type Team struct {
	name         string
	results      map[string]int // mapping for number of wins/draws/losses
	totalMatches int
	totalPoints  int
}

// Teams - a slice with elements of type team
type Teams []Team

// The outcome should be ordered by points, descending. In case of a tie, teams are ordered alphabetically
// Len - a method so we can sort
func (ts Teams) Len() int { return len(ts) }

// Less - a method so we can sort
func (ts Teams) Less(i, j int) bool {
	if ts[i].totalPoints == ts[j].totalPoints {
		return ts[i].name < ts[j].name // add alphabetic sort method
	}
	return ts[i].totalPoints > ts[j].totalPoints // descending order
}

// Swap - a method so we can sort a
func (ts Teams) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// GetTeam - a method to return the index of a specific team in the teams slice
func (ts Teams) GetTeam(name string) (i int, ok bool) {
	for i, v := range ts {
		if v.name == name {
			return i, true
		}
	}
	return -1, false
}

const (
	newline = "\n"
	linesep = ";"
)

var (
	scroreMap = map[string]int{"win": 3, "draw": 1, "loss": 0}
)

// Tally converts the outcomes list to a formatted results table which is available on an io.Writer interface
func Tally(r io.Reader, w io.Writer) error {
	b, err := ioutil.ReadAll(r) // just read all the bytes... could have used scanner.Scan
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), newline) // cast the bytes to a slice of strings
	var records []Team
	for _, line := range lines {
		line = strings.Trim(string(line), "\t\n")           // remove tabs and newline characters
		if len(line) <= 1 || strings.HasPrefix(line, "#") { // skip empty or commented lines
			continue
		}
		teamNames, scores, err := parseLine(line)
		if err != nil {
			return err
		}
		for i, name := range teamNames {
			if idx, ok := Teams(records).GetTeam(name); ok {
				// team found, update results
				records[idx].results[scores[i]]++
				records[idx].totalMatches++
				records[idx].totalPoints += scroreMap[scores[i]]
			} else {
				// team is not in list, add with result
				records = append(records,
					Team{name: name, results: map[string]int{scores[i]: 1}, totalMatches: 1, totalPoints: scroreMap[scores[i]]})
			}
		}
	}
	sort.Sort(Teams(records))
	w.Write([]byte(formatTable(records)))
	return nil
}

// parseLine - a helper to parse the line format given by the exercise
func parseLine(line string) (results []string, scores [2]string, err error) {
	results = strings.Split(line, linesep)
	if len(results) != 3 {
		return results, scores, errors.New("each line must contain 3 entries")
	}
	v, ok := scroreMap[results[2]]
	if !ok {
		return results, scores, errors.New("last line element must specify win, draw or loss")
	}
	switch v {
	case 0:
		scores = [2]string{"loss", "win"}
	case 1:
		scores = [2]string{"draw", "draw"}
	case 3:
		scores = [2]string{"win", "loss"}
	}
	return results[0:2], scores, nil
}

// formatTable - a helper to build the results table string
func formatTable(ts Teams) string {
	table := "Team                           | MP |  W |  D |  L |  P\n"
	for _, t := range ts {
		table += fmt.Sprintf("%-30v | % 2v | % 2v | % 2v | % 2v | % 2v\n",
			t.name, t.totalMatches, t.results["win"], t.results["draw"], t.results["loss"], t.totalPoints)
	}
	return table
}
