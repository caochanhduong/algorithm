func mySqrt(x int) int {
    if x == 1 {
        return 1
    }
    left, right := 0, x
    var mid int
    var sqr int
    for {
        mid = (left + right) / 2
        // case 1 stop: right is equal to left
        if mid == left {
            return mid
        }
        sqr = mid * mid
        // case 2 stop
        if sqr == x {
            return mid
        }
        if sqr < x {
            left = mid
        }
        if sqr > x {
            right = mid
        }
    }
}
