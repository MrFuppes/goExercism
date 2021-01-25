package perfect

import "errors"

// Classification - a type for Nicomachus' classification scheme for natural numbers
type Classification string

const (
	// ClassificationDeficient - n < sum
	ClassificationDeficient Classification = "ClassificationDeficient"
	// ClassificationPerfect - n == sum
	ClassificationPerfect Classification = "ClassificationPerfect"
	// ClassificationAbundant - n > sum
	ClassificationAbundant Classification = "ClassificationAbundant"
	// ErrOnlyPositive -

)

var (
	// ErrOnlyPositive - an error that shoud be defined
	ErrOnlyPositive = errors.New("n must be positive")
)

// Classify a function implementing the classification scheme
func Classify(n int64) (cls Classification, err error) {

	var sum int64

	if n < 1 {
		return "", ErrOnlyPositive
	}

	for i := n / 2; i > 0; i-- {
		if n%i == 0 {
			sum += i
		}
	}

	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	case sum < n:
		return ClassificationDeficient, nil
	}
	return
}
