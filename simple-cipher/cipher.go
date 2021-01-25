// Package cipher implements the shift cipher
package cipher

import "bytes"

// Cipher interface - types have to satisfy encoding and decoding functionality
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// Vigenere as a general implementation for shift cipher
type Vigenere struct {
	key string
}

// NewCaesar - a cipher interface with shift = +3
func NewCaesar() Cipher { return NewShift(3) }

// NewShift - a cipher with specified shift distance (-25 to +25, without 0)
func NewShift(distance int) Cipher {
	if distance < 0 {
		distance += 26
	}
	if distance > 25 || distance == 0 {
		return nil
	}

	return NewVigenere(string(rune('a' + distance)))
}

// NewVigenere - a cipher with specified shift distance for each character (-25 to +25, without 0)
func NewVigenere(key string) Cipher {
	var ok bool
	for _, c := range key {
		if c < 'a' || c > 'z' {
			return nil
		}
		if c != 'a' { // make sure at least one character is not 'a'
			ok = true
		}
	}
	if !ok {
		return nil
	}

	return &Vigenere{key}
}

// Encode a string using a given cipher key
func (cipher *Vigenere) Encode(s string) string {
	var buf bytes.Buffer
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			buf.WriteByte(byte(c) - 'A' + 'a')
		} else if c >= 'a' && c <= 'z' {
			buf.WriteByte(byte(c))
		}
	}

	s = buf.String()
	buf.Reset()
	for i, c := range s {
		buf.WriteByte('a' + ((byte(c)-'a')+(cipher.key[i%len(cipher.key)]-'a'))%26)
	}

	return buf.String()
}

// Decode a string using a given cipher key
func (cipher *Vigenere) Decode(s string) string {
	var buf bytes.Buffer
	for i, c := range s {
		buf.WriteByte('a' + ((byte(c)-'a')-(cipher.key[i%len(cipher.key)]-'a')+26)%26)
	}

	return buf.String()
}
