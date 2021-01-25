// Package react implements the react exercise
package react

// cell - basic building block of the table (a value-hoder).
type cell struct {
	reactor *table
	value   int
}

// table - the space holding the cells.
type table struct {
	cells []Cell
}

// New returns a pointer to a new table that constitutes the "reactor"
func New() Reactor {
	return &table{}
}

// CreateInput creates an input cell on the table space with the given initial value.
func (t *table) CreateInput(i int) InputCell {
	c := &cell{reactor: t, value: i}
	t.cells = append(t.cells, c)
	return c
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (t *table) CreateCompute1(c Cell, fn func(int) int) ComputeCell {
	cc := &computeCell{callbacks: make(map[*func(int)]func(int))}
	cc.reactor = t
	t.cells = append(t.cells, cc)

	cc.parents = []Cell{c}
	cc.compute = func(d ...Cell) int { return fn(d[0].Value()) }
	cc.value = cc.compute(c)
	return cc
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (t *table) CreateCompute2(c1, c2 Cell, fn func(int, int) int) ComputeCell {
	cc := &computeCell{callbacks: make(map[*func(int)]func(int))}
	cc.reactor = t
	t.cells = append(t.cells, cc)

	cc.parents = []Cell{c1, c2}
	cc.compute = func(parents ...Cell) int { return fn(parents[0].Value(), parents[1].Value()) }
	cc.value = cc.compute(c1, c2)
	return cc
}

// SetValue sets the value of a cell
func (c *cell) SetValue(i int) {
	if c.value != i {
		c.value = i
		c.updateCells()
	}
}

// Value returns the current value of a cell
func (c *cell) Value() int { return c.value }

// updateCells - a helper to update all cells of the table by calling thier value
func (c *cell) updateCells() {
	for _, v := range c.reactor.cells {
		v.Value()
	}
}

// computeCell - a cell that computes its value based on parent cell(s)
type computeCell struct {
	cell
	parents   []Cell
	compute   func(...Cell) int
	callbacks map[*func(int)]func(int)
}

// Value runs the compute function of a computedCell. If the new value differs from
// the previous, queued callback functions are executed.
func (cc *computeCell) Value() int {
	prev := cc.value
	cc.value = cc.compute(cc.parents...)
	if prev != cc.value {
		for _, f := range cc.callbacks {
			f(cc.value)
		}
	}
	return cc.value
}

// AddCallback adds a callback to a computeCell. Use callback function's address as unique key.
func (cc *computeCell) AddCallback(fn func(int)) Canceler {
	cc.callbacks[&fn] = fn
	return canceler(func() { delete(cc.callbacks, &fn) })
}

// canceler - we need a type that satisfies the Cancel interface
type canceler func()

// Cancel - the function needed for Cancel interface - just runs the canceler
func (c canceler) Cancel() { c() }
