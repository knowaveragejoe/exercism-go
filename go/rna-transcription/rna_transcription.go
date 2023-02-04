package strand

func ToRNA(dna string) string {
	rna := ""
	for _, n := range dna {
		switch n {
		case 'A':
			rna += string('U')
		case 'C':
			rna += string('G')
		case 'G':
			rna += string('C')
		case 'T':
			rna += string('A')
		}
	}

	return rna
}
