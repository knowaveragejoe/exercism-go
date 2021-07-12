package accumulate

func Accumulate(slice []string, converter func(string) string) []string {
	var output []string
	for _, word := range slice {
		output = append(output, converter(word))
	}

	return output
}
