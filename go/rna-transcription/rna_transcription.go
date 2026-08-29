package rnatranscription

func ToRNA(dna string) string {
	rna := []byte(dna)

	for i, nucleotide := range rna {
		switch nucleotide {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		}
	}

	return string(rna)
}
