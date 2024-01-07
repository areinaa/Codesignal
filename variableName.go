func solution(name string) bool {
    
    if (name[0]>=48 && name[0]<= 57){
        return false
    }
    
    for _, val := range name {
        if (val == 95 || (val >= 97 && val <=122) || (val >= 65 && val <=90) || (val>=48 && val<= 57)){
            continue
        }else{
            return false
        }
    }

    
    return true
}
