import "math"
func solution(inputArray []int) int {
    
    var maxAbs int
    var previousAbs int
    
    for i := 0; i < len(inputArray) - 1; i++ {
        maxAbs = int(math.Abs(float64(inputArray[i]) - float64(inputArray[i + 1])))
    
        if maxAbs > previousAbs {
            previousAbs = maxAbs
        }
    }

    return previousAbs
    
}
