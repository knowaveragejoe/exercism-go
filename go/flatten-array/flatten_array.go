package flatten

var flattened []interface{}

func Flatten(input interface{}) []interface{} {
	flattened := []interface{}{}
	DoFlatten(&flattened, input)

	return flattened
}

func DoFlatten(accum *[]interface{}, input interface{}) {
	// Assert the input slice is actually usable
	slice := input.([]interface{})

	for _, item := range slice {
		// Skip nil values
		if item == nil {
			continue
		}

		// If value is an int, append to accumulator
		switch t := item.(type) {
		case int:
			*accum = append(*accum, t)
		case []interface{}:
			//Otherwise, recurse
			DoFlatten(accum, t)
		}
	}
}