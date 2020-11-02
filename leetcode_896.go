func isMonotonic(A []int) bool {
    isIncrease, flag := false, false
    length := len(A)
    // use a flag to mark the first difference between two continuous elements in the array
    for i := 0; i < length - 1; i++ {
        if A[i+1] - A[i] != 0 {
            diff := A[i+1] - A[i]
            if !flag {
                flag = true
                if diff > 0 {
                    isIncrease = true
                }
                continue
            }
            if (isIncrease && diff < 0) || (!isIncrease && diff > 0) {
                return false   
            }
        }
    }
    return true
}

func isMonotonic(A []int) bool {
    isIncreasing, isDecreasing := true, true
    length := len(A)
    // use two boolean to check increase or decrease
    for i := 0; i < length - 1; i++ {
        if A[i+1] - A[i] > 0 {
            isDecreasing = false
            continue
        }
        if A[i+1] - A[i] < 0 {
            isIncreasing = false
        }
    }
    return isDecreasing || isIncreasing
}
