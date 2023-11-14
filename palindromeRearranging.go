func solution(inputString string) bool {
    
    chars := make(map[byte]int)
    
    for i:=0;i<len(inputString);i++{
        chars[inputString[i]]++
    }
    
    aux:=0
    for _, c := range chars{
        if c%2 != 0 && len(inputString)%2 == 0{
            return false
        }
        if c%2 != 0 && len(inputString)%2 != 0 && aux != 0{
            return false
        }
        if c%2 != 0 && len(inputString)%2 != 0 && aux == 0{
            aux++
        }
        
    }

    
    return true
}
