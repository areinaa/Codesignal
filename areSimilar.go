func solution(a []int, b []int) bool {
    aux:=0
    
    for i:=0;i<len(a);i++{
        if a[i]!=b[i] && aux == 0{
            for j:=i;j<len(b);j++{
                if a[i]==b[j] && aux == 0 && a[j]!=b[j]{
                    aux++
                    b[j]=b[i]
                    break
                }
                  
            }
            if aux != 1{
                return false
            }
             
        }else if a[i]!=b[i] && aux == 1{
            return false
        }  
    }
    return true
}
