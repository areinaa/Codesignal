func solution(sequence []int) bool {
    
    aux:=0
    for i:=0;i<len(sequence)-1;i++{
        if sequence[i]<sequence[i+1]{
            continue
        }else if i+2==len(sequence) && aux==0{
            aux++
        }else if i==0 || (sequence [i-1] < sequence[i+1] && aux == 0){ 
            aux++
        }else if i+2<len(sequence) && sequence[i]<sequence[i+2] && aux==0{
            aux++
        }else{
            return false
        }   
    }
    return true
}
