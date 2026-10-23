func hasDuplicate(nums []int) bool {
    mapNums := make(map[int]int, len(nums))
    for i, num := range nums {
        if _ , x := mapNums[num]; x {
            return true
        }
        mapNums[num] = i
    }
    return false
}
