package matrix

func UpdateMatrixDP(mat [][]int) [][]int {
    maxRows := len(mat)
    maxCols := len(mat[0])
    
    distances := mat

    // update columns
    for rowIdx := 0; rowIdx < maxRows; rowIdx++ {
	for colIdx := 1; colIdx < maxCols; colIdx++ {
		if mat[rowIdx][colIdx] != 0 {
			distances[rowIdx][colIdx] = distances[rowIdx][colIdx-1] + 1
		}
	}
    }
    for rowIdx := 0; rowIdx < maxRows; rowIdx++ {
	for colIdx := 0; colIdx < maxCols-1; colIdx++ {
		if distances[rowIdx][colIdx] > distances[rowIdx][colIdx+1] + 1 {
			distances[rowIdx][colIdx] = distances[rowIdx][colIdx+1] + 1
		}
	}
    }
 
    // update rows
    for rowIdx := 1; rowIdx < maxRows; rowIdx++ {
	for colIdx := 0; colIdx < maxCols; colIdx++ {
		if distances[rowIdx][colIdx] > distances[rowIdx-1][colIdx] + 1 {
			distances[rowIdx][colIdx] = distances[rowIdx-1][colIdx] + 1
		}
	}
    }
    for rowIdx := 0; rowIdx < maxRows-1; rowIdx++ {
	for colIdx := 0; colIdx < maxCols; colIdx++ {
		if distances[rowIdx][colIdx] > distances[rowIdx+1][colIdx] + 1 {
			distances[rowIdx][colIdx] = distances[rowIdx+1][colIdx] + 1
		}
	}
    }

    return distances
}
