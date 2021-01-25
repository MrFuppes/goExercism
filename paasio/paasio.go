package paasio

import (
	"io"
	"sync"
)

// counter - a struct to store results.
// also provides the reader/writer functions together with a mutex to make the
// operation concurrency-safe.
type counter struct {
	n         int64
	nops      int
	readFunc  func(p []byte) (n int, err error)
	writeFunc func(p []byte) (n int, err error)
	mutex     sync.Mutex
}

// ReadCount returns the number of bytes read and number of operations from counter
func (c *counter) ReadCount() (n int64, nops int) {
	return c.n, c.nops
}

// WriteCount returns the number of bytes written and number of operations from counter
func (c *counter) WriteCount() (n int64, nops int) {
	return c.n, c.nops
}

// Read performs a read operation and adds the number of bytes read / number of operations to the counter.
// During the read operation, the counter is locked by the mutex so that no other function can modifiy it simultaneously.
func (c *counter) Read(p []byte) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	n, err := c.readFunc(p)
	if err == nil {
		c.n += int64(n)
		c.nops++
	}
	return n, err
}

// Write performs a write operation and adds the number of bytes written / number of operations
// During the write operation, the counter is locked by the mutex so that no other function can modifiy it simultaneously.
func (c *counter) Write(p []byte) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	n, err := c.writeFunc(p)
	if err == nil {
		c.n += int64(n)
		c.nops++
	}
	return n, err
}

// NewReadCounter returns the address of a zeroed counter with a defined read function.
func NewReadCounter(r io.Reader) ReadCounter {
	return &counter{n: 0, nops: 0, readFunc: r.Read, writeFunc: nil}
}

// NewWriteCounter returns the address of a zeroed counter with a defined write function.
func NewWriteCounter(w io.Writer) WriteCounter {
	return &counter{n: 0, nops: 0, writeFunc: w.Write, readFunc: nil}
}

// NewReadWriteCounter  the address of a zeroed counter with a defined read and write function.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &counter{n: 0, nops: 0, readFunc: rw.Read, writeFunc: rw.Write}
}
