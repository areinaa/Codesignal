func solution(picture []string) []string {

    var border string
    sol:= []string{}
    
    for i:=0;i<len(picture[0])+2;i++{
        border +="*"
    }
    sol=append(sol,border)
    
    for i:=0;i<len(picture);i++{
        sol=append(sol,"*" + picture[i] + "*")
    }
    sol=append(sol,border)

    return sol
}
