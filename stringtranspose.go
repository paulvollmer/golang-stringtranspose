package stringtranspose

// Transpose process a two dimensional array of the same size
func Transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

// TransposeIrregular process a two dimensional array of an irregular size
func TransposeIrregular(slice [][]string) [][]string {
	// find the longest array
	longest := 0
	xLenghts := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		if longest < len(slice[i]) {
			longest = len(slice[i])
		}
		xLenghts[i] = len(slice[i])
	}

	nslice := make([][]string, len(slice))
	for i := 0; i < len(slice); i++ {
		nslice[i] = make([]string, longest)
		for j := 0; j < len(slice[i]); j++ {
			nslice[i][j] = slice[i][j]
		}
	}

	return Transpose(nslice)
}
