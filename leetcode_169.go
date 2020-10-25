func majorityElement(nums []int) int {
    mapNum := make(map[int]int)
    fmt.Println(mapNum[0])
    n := len(nums)
    for i := 0; i < n; i++ {
        mapNum[nums[i]]++
        if mapNum[nums[i]] > n/2 {
            return nums[i]
        }
    }
    return -1
}
