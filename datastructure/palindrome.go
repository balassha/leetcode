func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x == 0 {
        return true
    }
    var arr []int
    for x != 0 {
        arr = append(arr,x%10)
        x = x/10
    }
    mid := 0
    if len(arr)%2 == 0 {
        mid = len(arr)/2
    }else{
        mid = len(arr)/2-1
    }
    for i:=0;i<=mid;i++{
            if arr[i]==arr[len(arr)-1-i]{
                continue
            }else{
                return false
            }
    }
    return true
}
