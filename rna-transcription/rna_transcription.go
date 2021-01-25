package strand

// DNAtoRNA - a mapping of DNA nucleotides to RNA nucleotides
var DNAtoRNA = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA - convert DNA sequence to corresponding RNA seq.
func ToRNA(dna string) string {
	if len(dna) == 0 {
		return ""
	}

	var rna string
	for _, r := range dna {
		rna += DNAtoRNA[r]
	}
	return rna
}
