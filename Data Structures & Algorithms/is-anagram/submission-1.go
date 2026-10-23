func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var counts [26]int

    // Итерируемся по байтам без декодирования UTF-8
    for i := 0; i < len(s); i++ {
        counts[s[i]-'a']++
        counts[t[i]-'a']--
    }

    for _, c := range counts {
        if c != 0 {
            return false
        }
    }

    return true
}