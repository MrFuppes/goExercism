package rectangles

const (
	corner = '+'
	wall   = '|'
	ceil   = '-'
	space  = ' '
)

// Count returns the number of rectangles displayed in input slice
func Count(input []string) (count int) {
	h := len(input)
	if h <= 1 {
		return count
	}
	w := len(input[0])
	for y := 0; y < h-1; y++ {
		for x := 0; x < w-1; x++ {
			if input[y][x] == corner {
				count += findRect(input, x, y, w, h)
			}
		}
	}
	return count
}

func findRect(input []string, x0, y0, w, h int) int {
	count := 0
	for i := y0 + 1; i < h; i++ {
		for j := x0 + 1; j < w; j++ {
			if input[i][j] == corner {
				count += checkRectangle(input, x0, y0, j, i)
			}
		}
	}
	return count
}

func checkRectangle(input []string, x0, y0, x1, y1 int) int {
	for i := y0; i <= y1; i++ {
		if input[i][x0] == space || input[i][x0] == ceil || input[i][x1] == space || input[i][x1] == ceil {
			return 0
		}
	}
	for j := x0; j <= x1; j++ {
		if input[y0][j] == space || input[y0][j] == wall || input[y1][j] == space || input[y1][j] == wall {
			return 0
		}
	}
	return 1
}
