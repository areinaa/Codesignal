func ordenarFun(s string) string{
   
    fin := []string{}
    
    for i:=0;i<len(s);i++{
        fin = append(fin,string(s[len(s)-1-i]))
    }
    

    return strings.Join(fin,"")
}

func solution(inputString string) string {
  
    pos := []int{}

    for i:=0;i<len(inputString);i++{
        c := inputString[i]
        if c == '('{

            pos = append(pos, i)
        }else if c == ')'{
            n := pos[len(pos)-1]
            pos = pos[:len(pos)-1]
        
            reordenar := ordenarFun(inputString[n+1:i])
            inputString = inputString[:n] + reordenar + inputString[i+1:]
            
            i-=2 
        }
        
    }
    
    return inputString
}
