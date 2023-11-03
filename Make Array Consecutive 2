func solution(statues []int) int {
       
    result := 1
    min := statues[0]
    max := statues [0]
    
    for _, value := range statues{
        if value <= min {
            result += min - value -1
            min = value
            fmt.Println(result)
        }else if value > max {
            result += value - max - 1
            max = value
            fmt.Println(result)
        }else{
            result--
            fmt.Println(result)
        }
    }
    return result
}
