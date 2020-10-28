func containsNearbyDuplicate(nums []int, k int) bool {
    mapNum := make(map[int]int)
    
    for index, item := range nums {
        if v, ok := mapNum[item]; ok && index - v <= k {
            return true
        }
        mapNum[item] = index
    }
    
    return false
}
