func solution(image [][]int) [][]int {

    var result [][]int
    
    for i := 0; i<len(image)-2 ;i++{
        var slice []int
        for j := 0; j<len(image[1])-2;j++{
            value := image[i][j]+image[i][j+1]+image[i][j+2]+image[i+1][j]+image[i+1][j+1]+image[i+1][j+2]+image[i+2][j]+image[i+2][j+1]+image[i+2][j+2]
            value = value/9
            
            slice = append(slice, value)
        }
        result = append(result, slice)
    }
    

return result
}
