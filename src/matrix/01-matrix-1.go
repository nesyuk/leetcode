package matrix
import (
	"math"
)

func UpdateMatrixBST(mat [][]int) [][]int {
    maxRows := len(mat)
    maxCols := len(mat[0])
    
    distances := make([][]int, maxRows)
    for i := 0; i < maxRows; i++ {
        distances[i] = make([]int, maxCols)
        
    }
    
    queue := [][]int{}
    for rowIdx := 0; rowIdx < maxRows; rowIdx++ {
        for colIdx := 0; colIdx < maxCols; colIdx++ {
            if mat[rowIdx][colIdx] == 0 {
                queue = append(queue, []int{rowIdx, colIdx})
            } else {
                distances[rowIdx][colIdx] = math.MaxInt64
            }
        }
    }
    directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
    for len(queue) != 0 {
        cell := queue[0]
        queue = queue[1:]
        for _, direction := range directions {
            adjRow := cell[0] + direction[0]
            adjCol := cell[1] + direction[1]
            if adjRow >= 0 && adjRow < maxRows && adjCol >= 0 && adjCol < maxCols {
                if distances[adjRow][adjCol] > distances[cell[0]][cell[1]] + 1 {
                    distances[adjRow][adjCol] = distances[cell[0]][cell[1]] + 1
                    queue = append(queue, []int{adjRow, adjCol})
                }
            }
        }
    }
    
    return distances
}
