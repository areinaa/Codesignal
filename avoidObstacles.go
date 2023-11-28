func solution(inputArray []int) int {
    
    sol:=1
    
    for i:=0;i<len(inputArray);i++{
        if inputArray[i]%sol == 0{
            sol++
            i=-1
        }
    }
    
    return sol
}
