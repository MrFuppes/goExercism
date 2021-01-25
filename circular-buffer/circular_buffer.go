package circular

import "errors"

var (
	errIsEmpty = errors.New("buffer is empty")
	errIsFull  = errors.New("buffer is full")
)

// Buffer to (ring-)buffer reads and writes of bytes
type Buffer struct {
	buf      []byte
	size     int
	occupied int
	r        int // pointer to read position
	w        int // pointer to write position
}

// NewBuffer returns pointer to a ring buffer with a defined size
func NewBuffer(size int) *Buffer {
	return &Buffer{buf: make([]byte, size), size: size}
}

// ReadByte reads a byte from the buffer
func (b *Buffer) ReadByte() (byte, error) {
	if b.occupied == 0 {
		return 0, errIsEmpty
	}

	defer func(b *Buffer) { // inc read index AFTER returning the byte
		b.r++
		if b.r == b.size {
			b.r = 0
		}
		b.occupied--
	}(b)

	return b.buf[b.r], nil
}

// WriteByte writes a byte to the buffer
func (b *Buffer) WriteByte(c byte) error {
	if b.occupied == b.size {
		return errIsFull
	}
	b.buf[b.w] = c
	b.w++
	if b.w == b.size {
		b.w = 0
	}
	b.occupied++

	return nil
}

// Overwrite overwrits a byte
func (b *Buffer) Overwrite(c byte) {
	if b.occupied < b.size {
		b.WriteByte(c) // call normal write if buffer isn't full
	} else { // otherwise overwrite, then inc both read and write pointer
		b.buf[b.w] = c
		b.w++
		if b.w == b.size {
			b.w = 0
		}
		b.r++
		if b.r == b.size {
			b.r = 0
		}
	}
}

// Reset puts buffer in an empty state
func (b *Buffer) Reset() {
	b.buf = make([]byte, b.size)
	b.occupied, b.r, b.w = 0, 0, 0
}
