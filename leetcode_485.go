func findMaxConsecutiveOnes(nums []int) int {
    res, tmp := 0, 0
    for i := range nums {
        if nums[i] == 1 {
            tmp++
        } else {
            res = max(res, tmp)
            tmp = 0
        }
    }
    return max(res, tmp)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
