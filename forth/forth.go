// Package forth implements a modified version of DroidLinq's solution
package forth

import (
	"errors"
	"strconv"
	"strings"
)

var (
	operations = map[string]interface{}{
		"+":    add,
		"-":    subtract,
		"*":    multiply,
		"/":    divide,
		"dup":  duplicate,
		"drop": drop,
		"swap": swap,
		"over": over,
	}
	minStackLen = map[string]int{
		"+":    2,
		"-":    2,
		"*":    2,
		"/":    2,
		"dup":  1,
		"drop": 1,
		"swap": 2,
		"over": 2,
	}
	errInsufficientParms = errors.New("insufficient parameters")
	errNbrRedefinition   = errors.New("redefinition of number")
	errUnknownOperation  = errors.New("unknown operation")
	errZeroDiv           = errors.New("division by zero")
)

// Forth evaluator
func Forth(inputs []string) ([]int, error) {
	stack := []int{}
	userDefinedOps := map[string][]string{} // user-defined -> built-in

	for _, input := range inputs {
		if strings.HasPrefix(input, ":") && strings.HasSuffix(input, ";") {
			if err := parseUserDefinedOp(input, &userDefinedOps); err != nil {
				return nil, err
			}
		} else {
			return runOps(stack, strings.Fields(input), userDefinedOps)
		}
	}

	return nil, nil
}

// parseUserDefinedOp parses a user-defined operation to built-in operations
func parseUserDefinedOp(input string, userDefinedOps *map[string][]string) error {
	elements := strings.Fields(input)
	if _, err := strconv.Atoi(elements[1]); err == nil {
		return errNbrRedefinition
	}

	ops := []string{} // collect the operations as built-in ops
	for _, op := range elements[2 : len(elements)-1] {
		if bulitInOp, ok := (*userDefinedOps)[strings.ToLower(op)]; ok {
			ops = append(ops, bulitInOp...)
		} else {
			ops = append(ops, op)
		}
	}
	(*userDefinedOps)[strings.ToLower(elements[1])] = ops

	return nil
}

func runOps(stack []int, ops []string, userDefinedOps map[string][]string) ([]int, error) {
	for _, op := range ops {
		if n, err := strconv.Atoi(op); err == nil { // numbers only get pushed onto stack
			stack = append(stack, n)
		} else if userOp, ok := userDefinedOps[strings.ToLower(op)]; ok { // re-call function with user-defined operation
			if stack, err = runOps(stack, userOp, userDefinedOps); err != nil {
				return stack, err
			}
		} else if f, ok := operations[strings.ToLower(op)]; ok {
			if len(stack) < minStackLen[op] { // assert sufficient parameters
				return stack, errInsufficientParms
			}
			if op == "/" && stack[len(stack)-1] == 0 { // prevend divide by zero
				return stack, errZeroDiv
			}
			stack = f.(func([]int) []int)(stack)
		} else {
			return stack, errUnknownOperation
		}
	}

	return stack, nil
}

func add(stack []int) []int {
	stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
	return stack[:len(stack)-1]
}

func subtract(stack []int) []int {
	stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
	return stack[:len(stack)-1]
}

func multiply(stack []int) []int {
	stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
	return stack[:len(stack)-1]
}

func divide(stack []int) []int {
	stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
	return stack[:len(stack)-1]
}

func duplicate(stack []int) []int {
	return append(stack, stack[len(stack)-1])
}

func drop(stack []int) []int {
	return stack[0 : len(stack)-1]
}

func swap(stack []int) []int {
	stack[len(stack)-2], stack[len(stack)-1] = stack[len(stack)-1], stack[len(stack)-2]
	return stack
}

func over(stack []int) []int {
	return append(stack, stack[len(stack)-2])
}
