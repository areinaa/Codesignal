func solution(n int) bool {
    
    num := strconv.Itoa(n)
    
    for i:=1 ; i<=len(num) ; i++ {
        a,_ := strconv.Atoi(num[i-1:i])
        if  a % 2 == 0{
            continue
        }else{
            return false
        }
         
    }

    return true
}
