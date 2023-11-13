func solution(a []int) int {
    
     moves := 0
    for i := 1; i < len(a); i++ {
        if d := a[i-1] - a[i]; d >= 0 {
            a[i] += d + 1
            moves += d + 1
        }
    }
    
    return moves   
}
