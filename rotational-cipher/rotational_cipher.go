package rotationalcipher

import "strings"

// RotationalCipher implements a simple rotation around the alphabet
func RotationalCipher(plain string, key int) string {
	var encoded strings.Builder

	for _, c := range plain {
		if c >= 'A' && c <= 'Z' {
			encoded.WriteByte((byte(c)-'A'+byte(key))%26 + 'A')
		} else if c >= 'a' && c <= 'z' {
			encoded.WriteByte((byte(c)-'a'+byte(key))%26 + 'a')
		} else {
			encoded.WriteByte(byte(c))
		}
	}

	return encoded.String()
}
