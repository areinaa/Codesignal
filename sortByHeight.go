import "sort"
func solution(a []int) []int {
  
  personIndices := make([]int, 0)
  personHeights := make([]int, 0)
    
    for i, h := range a {
        if h != -1 {
            personIndices = append(personIndices, i)
            personHeights = append(personHeights, h)
        }
    }
    
    sort.Ints(personHeights)
    
    for j, index := range personIndices {
        a[index] = personHeights[j]
    }
    
    return a
    
}
