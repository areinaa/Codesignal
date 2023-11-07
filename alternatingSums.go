func solution(a []int) []int {
    
    sol := []int{0,0}
    
    for i, n := range a{
        
        if i%2 ==0 {
            sol[0]+=n
        }else{
            sol[1]+=n
        }
        
    }
    return sol
}
