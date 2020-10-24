func summaryRanges(nums []int) []string {
    n := len(nums)
    res := []string{}
    if n == 0 {
        return []string{}
    }
    
    if n == 1 {
        return []string{strconv.Itoa(nums[0])}
    }
    
    tmp := nums[0]
    
    for i := 0; i < n; i++ {
        if i < n - 1 && nums[i] + 1 != nums[i+1] || i == n - 1 {
            if nums[i] != tmp {
                res = append(res, fmt.Sprintf("%d->%d", tmp, nums[i]))
                if i < n - 1 {
                    tmp = nums[i+1]
                }
                continue
            }
            res = append(res, strconv.Itoa(tmp))
            if i < n - 1 {
                tmp = nums[i+1]
            }
        }
    }
    return res
}
