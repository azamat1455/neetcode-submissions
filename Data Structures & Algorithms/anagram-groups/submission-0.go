func groupAnagrams(strs []string) [][]string {
	strsMap := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		currStr := strs[i]
		var currArr [26]int
		for j := 0; j < len(currStr); j++ {
			currArr[currStr[j] - 'a'] ++
		}
		strsMap[currArr] = append(strsMap[currArr], currStr)
	}

	result := make([][]string, 0, len(strsMap))
    for _, group := range strsMap {
        result = append(result, group)
    }

    return result
}
