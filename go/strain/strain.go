package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var filtered []int

	for _, n := range i {
		if filter(n) {
			filtered = append(filtered, n)
		}
	}

	return filtered
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var filtered []int

	for _, n := range i {
		if !filter(n) {
			filtered = append(filtered, n)
		}
	}

	return filtered
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var filtered [][]int

	for _, n := range l {
		if filter(n) {
			filtered = append(filtered, n)
		}
	}

	return filtered
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var filtered []string

	for _, n := range s {
		if filter(n) {
			filtered = append(filtered, n)
		}
	}

	return filtered
}
