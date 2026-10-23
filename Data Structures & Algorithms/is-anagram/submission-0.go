func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapR := make(map[rune]int, len(s))

    for _, r := range s {
        mapR[r]++
    }

    for _, r := range t {
        mapR[r]--

        if mapR[r] < 0 {
            return false
        }
    }

    return true
}