// Package variablelengthquantity implements VLQ
// Wikipedia: A VLQ is essentially a base-128 representation of an unsigned integer with the addition of the
// eighth bit to mark continuation of bytes
package variablelengthquantity

import "errors"

// EncodeVarint encodes 32 bit uints to a vlq byte sequence
func EncodeVarint(seq []uint32) (encoded []byte) {
	for _, v := range seq {
		tmp := []byte{}
		for v > 0x7f {
			// to byte, set continuation bit and prepend
			tmp = append([]byte{byte(v) | 0x80}, tmp...)
			// next 7 bits...
			v >>= 7
		}

		// v < 0x7f remaining
		tmp = append([]byte{byte(v) | 0x80}, tmp...)

		// first byte could be greater 128 (== invalid)
		if tmp[len(tmp)-1] >= 0x80 {
			tmp[len(tmp)-1] -= 0x80
		}
		encoded = append(encoded, tmp...)
	}

	return encoded
}

// DecodeVarint decodes a vlq byte sequence to a sequence of uints
func DecodeVarint(vlq []byte) (decoded []uint32, err error) {
	tmp := uint32(0) // could also use uint64 here
	for _, b := range vlq {
		if b < 0x80 { // continuation bit not set
			tmp |= uint32(b)
			decoded = append(decoded, tmp)
			tmp = 0
			continue // to next byte sequence
		}
		b -= 0x80 // remove continuation bit, same as &^ 128
		tmp |= uint32(b)
		tmp <<= 7 // move bits to the left to make room for the next
	}
	if decoded == nil {
		return decoded, errors.New("incomplete sequence")
	}

	return decoded, err
}
