func plusOne(digits []int) []int {
    n := len(digits)
    for i := n - 1; i >= 0; i-- {
    	digits[i] += 1
        
      	if digits[i] <= 9 {
        	return digits
      	}
        
        digits[i] = 0
        
        if i == 0 {
          	return append([]int{1}, digits...)
        }
    }
    return digits
}
