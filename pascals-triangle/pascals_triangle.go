package pascal

// Triangle computes n number of rows of Pascal's triangle
func Triangle(n int) (triangle [][]int) {
	for y := 1; y <= n; y++ { // y: row
		triangle = append(triangle, []int{}) // prepare a slice for the current row
		v := 1                               // first value is always 1
		for x := 1; x <= y; x++ {            // x: row element
			triangle[y-1] = append(triangle[y-1], v)
			v = v * (y - x) / x // calculate next value
		}
	}
	return triangle
}
