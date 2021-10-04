// https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
    var maxL, start int = 0, 0

    for i := 0; i <= len(s)-1; i++ {
        odd := getLen(s, i, i)
        even := getLen(s, i, i + 1)
        
        cur := 0
        if (odd >= even) {
            cur = odd
        } else {
            cur = even
        }

        if (cur > maxL) {
            maxL = cur
            start = i - (cur - 1) / 2
        }
    }

    // fmt.Println(start, maxL)
    return s[start:start + maxL]
}

func getLen(s string, l int, r int) int {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l -= 1
        r += 1
    }

    return r - l - 1
}
