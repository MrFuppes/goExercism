package allyourbase

import "errors"

// ConvertToBase converts the number in input base to output base
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (output []int, err error) {
	if inputBase < 2 {
		return []int{0}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{0}, errors.New("output base must be >= 2")
	}
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}
	if len(inputDigits) == 1 && inputDigits[0] == 0 {
		return []int{0}, nil
	}

	var decimal int
	for _, v := range inputDigits {
		if v < 0 || v >= inputBase {
			return output, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decimal = decimal*inputBase + v
	}

	if decimal == 0 {
		return []int{0}, nil
	}

	for decimal > 0 {
		output = append([]int{decimal % outputBase}, output...)
		decimal /= outputBase
	}

	return output, err
}
