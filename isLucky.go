func solution(n int) bool {
    
    strin := strconv.FormatInt(int64(n),10)
    result := 0
        
    for i:=0;i<len(strin)/2;i++{
        result += int(strin[i]) - int(strin[len(strin)-1-i])
    }
    
    if result == 0 {return true}
    return false

}
