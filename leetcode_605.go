func canPlaceFlowers(flowerbed []int, n int) bool {
    length := len(flowerbed)
    if length == 1 && flowerbed[0] == 0 || n == 0 {
        return true
    }
    
    for i := 0; i < length; i++ {
        if flowerbed[i] == 0 && ((i == 0  && i < length - 1 && flowerbed[i+1] == 0) || (i == length - 1 && i - 1 >= 0 && flowerbed[i-1] == 0) || (i - 1 >= 0 && i < length - 1 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0)) {
            flowerbed[i] = 1
            n--
            if n == 0 {
                return true
            }
        }
    }
        
    return false
}
