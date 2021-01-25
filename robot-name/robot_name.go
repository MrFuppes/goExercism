package robotname

import (
	"errors"
	"math/rand"
	"time"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums    = "0123456789"
)

var (
	allnames = map[string]int{}
	seed     *rand.Rand
)

// init - here, initialize the random seed
func init() {
	seed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// randStringFromChars - a helper to generate a random robot name from 2 letters and 3 numbers
func randStringFromChars(n int, chars string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[seed.Intn(len(chars))]
	}
	return string(b)
}

// Robot - as struct to represent a robot. For now, it only has a name property...
type Robot struct {
	name string
}

// generateName - generate a name an make sure it hasn't been used before
func generateName() (name string, err error) {
	if len(allnames) == 26*26*10*10*10 {
		return "", errors.New("out of names")
	}
	name = randStringFromChars(2, letters) + randStringFromChars(3, nums)
	allnames[name]++
	if allnames[name] > 1 { // recursive call to generateName if name already existed
		name, err = generateName()
	}
	return name, err
}

// Name - assigns a robot name if it doesn't have one yet
func (r *Robot) Name() (name string, err error) {
	name = r.name
	if name == "" {
		name, err = generateName()
		r.name = name
	}
	return name, err
}

// Reset - assign a new random name to an existing robot
func (r *Robot) Reset() (err error) {
	name, err := generateName()
	r.name = name
	return err
}
