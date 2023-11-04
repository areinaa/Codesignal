func solution(s1 string, s2 string) int {
    
    m := make(map[rune]int)
    result := 0
    
    for _, char := range s1{
        m[char]++
    }
    for _, char := range s2{
        if m[char] > 0{
            result++
            m[char]--
        }
    }
 
    return result
}
