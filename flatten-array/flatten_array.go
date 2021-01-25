// Package flatten contains functionality to flatten slices of things (interfaces).
package flatten

// flattener - a helper function to recursively flatten slices of "things"
func flattener(args []interface{}, v interface{}) []interface{} {
	// check if the passed thing "v" is slice-like with type assertion:
	if s, ok := v.([]interface{}); ok {
		for _, v := range s { // recursive call to elements of the slice
			args = flattener(args, v)
		}
	} else {
		if v != nil { // only append if value is not nil
			args = append(args, v)
		}
	}
	return args
}

// Flatten flattens a nested structure of arbitrary depth and arbitrary dtype
func Flatten(nested interface{}) []interface{} {
	return flattener([]interface{}{}, nested)
}
