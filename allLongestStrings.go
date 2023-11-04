func solution(inputArray []string) []string {

    result := []string{}
    maxlen := 0
    
    for _, value := range inputArray {
        if len(value)== maxlen {
            result = append(result, value)
        }else if len(value)>maxlen{
            maxlen = len(value)
            result = nil
            result = append(result,value)
        }
    }
    return result
}
