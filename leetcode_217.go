func containsDuplicate(nums []int) bool {
    mapNum := make(map[int]bool)
    for _, item := range nums {
        if _, ok := mapNum[item]; ok {
            return true
        }
        
        mapNum[item] = true
    }
    
    return false
}
