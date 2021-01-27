package challenge_go

func IterativeFactorial(nb int) int {
    result := 1
    if nb > 0 {
        for i := 1; i <= nb; i++ {
            result *= i
        }
    } else if nb == 0 {
        result = 1
    } else {
        result = 0
    }
    return result
}