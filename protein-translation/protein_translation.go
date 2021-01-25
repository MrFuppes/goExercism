// Package protein holds functions and errors to work with RNA strings
package protein

import (
	"errors"
)

// ErrStop - an error to be raised if a STOP base is encountered
var ErrStop = errors.New("stop sequence encountered")

// ErrInvalidBase - an error to be raised if an invalid lettercode is encountered
var ErrInvalidBase = errors.New("invalid base")

// translationMap - a mapping of lettercodes to bases
var translationMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

// FromCodon translates a lettercode to a base name
func FromCodon(codon string) (string, error) {
	base, ok := translationMap[codon]
	if ok {
		if base == "STOP" {
			return "", ErrStop
		}
		return base, nil
	}
	return "", ErrInvalidBase
}

// FromRNA translates a sequence of lettercodes to a sequence of protein names.
// Will return immediately if a STOP protein or an invalid base is encountered.
func FromRNA(rna string) ([]string, error) {
	var result []string
	for i := 0; i <= len(rna)-3; i += 3 {
		base, err := FromCodon(rna[i : i+3])
		if err == ErrInvalidBase {
			return result, err
		}
		if err == ErrStop {
			return result, nil
		}
		result = append(result, base)
	}
	return result, nil
}
