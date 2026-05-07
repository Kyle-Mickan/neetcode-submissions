func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    counts := make(map[rune]int)
    for _, ch := range s {
        counts[ch]++
    }
    for _, ch := range t {
        counts[ch]--
        if counts[ch] < 0 {
            return false
        }
    }
    return true
}
