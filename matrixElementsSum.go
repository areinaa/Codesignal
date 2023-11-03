func solution(matrix [][]int) int {
   
    totalPrice:=0
    
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if matrix[i][j]==0{
                for k:=i;k<len(matrix);k++{
                    matrix[k][j]=0
                }
            }else{
                totalPrice += matrix[i][j]
            }
            
        }
         
    }
    
    return totalPrice
}
