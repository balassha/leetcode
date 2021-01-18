func reverse(x int) int {
    var n int
    for x != 0 {
        n = n*10 + x%10
        x = x/10   
    }
    if n > 0 {
        if n > math.MaxInt32{
            return 0
        }
    }else{
        if n*-1 > math.MaxInt32{
            return 0
        }
    }
    return n
}
