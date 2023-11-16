func solution(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
    
    if yourLeft+yourRight == friendsLeft+friendsRight && (yourLeft==friendsLeft||yourRight==friendsLeft||yourLeft==friendsRight){
        return true
    }else{
        return false
    }

}
