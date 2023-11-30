func solution(matrix [][]bool) [][]int {
    
    result := make([][]int, len(matrix))
    
    for i, arr := range matrix {
        result[i] = make([]int, len(arr))
        
        for j, _ := range arr {
            
            for k:=i-1;k<=i+1;k++{
                
                if k < 0 || k >= len(matrix) {
                    continue
                }
                for z:=j-1;z<=j+1;z++{
                    
                    if z < 0 || (k == i && z == j) || z >= len(arr) {
                        continue
                    }
                    if matrix[k][z] {
                        result[i][j]++
                    }
                }
            }
        }
    }
    return result
}
