import "net"
func solution(inputString string) bool {
    return net.ParseIP(inputString) != nil
   /* 
    ip := strings.Split(inputString,".")

    fmt.Println(ip)
    if len(ip)!=4{
        return false
    }

    for _, n:=range ip{
        num,ok := strconv.Atoi(n)
        if ok != nil {
            return false
        }
        if n == "" {
            return false
        } else if num > 255 || num < 0{
            return false
        }
        
        
    }
    
    return true
    */
}
