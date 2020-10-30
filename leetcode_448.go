func findDisappearedNumbers(nums []int) []int {
    // init slice with capacity to reduce the extend action when appending to it later
    res := make([]int, 0, len(nums))
    
    // mark the element at index "num" with negative value, but still consider the absolute value of num
    for _, num := range nums {
        if nums[int(math.Abs(float64(num))) - 1] > 0 {
            nums[int(math.Abs(float64(num))) - 1] *= -1
        }
    }
    
    for i, num := range nums {
        if num > 0 {
            res = append(res, i+1)
        }
    }
    
    
    return res
}
