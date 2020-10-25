func singleNumber(nums []int) int {
    a := 0
    // special feature: b xor 0 = b, the elements which appear two times will produce zero, the element that appears only one time will remain
    for i := 0; i < len(nums); i++ {
        a ^= nums[i]
    }
    return a
}
