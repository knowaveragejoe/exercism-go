package protein

import (
	"errors"
)

var ErrStop = errors.New("Stop detected")
var ErrInvalidBase = errors.New("Invalid base detected")

func FromRNA(rna string) ([]string, error) {
	rnaRunes := []rune(rna)
	var codons []string

	for i, _ := range rnaRunes {
		if i > 0 && (i+1)%3 == 0 {
			codon, err := FromCodon(string(rnaRunes[i-2 : i+1]))
			if codon != "" {
				codons = append(codons, codon)
			} else {
				if err == ErrStop {
					return codons, nil
				} else {
					return codons, err
				}
			}
		}
	}

	return codons, nil
}

func FromCodon(codon string) (string, error) {
	if len(codon) != 3 {
		return "", ErrInvalidBase
	}

	proteins := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	protein := proteins[codon]
	if protein != "" {
		if protein == "STOP" {
			return "", ErrStop
		}
		return protein, nil
	}

	return "", ErrInvalidBase
}
