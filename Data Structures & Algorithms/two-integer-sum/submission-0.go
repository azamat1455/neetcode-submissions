func twoSum(nums []int, target int) []int {
    mapNum := make(map[int]int, len(nums))
    
    for i, num := range nums {
        complement := target - num

        if v, ok := mapNum[complement]; ok {
            return []int{v, i}
        }
        
        mapNum[num] = i
    }

    return nil
}