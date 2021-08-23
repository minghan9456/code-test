// https://leetcode.com/problems/count-sorted-vowel-strings/

/*
Math
H 5 1
(5 + n -1)! / n! * (5 - 1)! = (n + 4)! / n! * 4!
= (n + 1) * (n + 2) * (n + 3) * (n + 4) / 24

func countVowelStrings(n int) int {
    return ((n + 1) * (n + 2) * (n + 3) * (n + 4)) / 24
}
*/ 

// DP
func countVowelStrings(n int) int {
    a, e, i, o, u := 1, 1, 1, 1, 1
    for idx := 0; idx < n - 1; idx++ {
        // add new char before prev string
        a = a + e + i + o + u // a, e, i, o, u -> aa, ae, ai, ao, au
        e = e + i + o + u // e, i, o, u -> ee, ei, eo, eu
        i = i + o + u // i, o, u -> ii, io, iu
        o = o + u // o, u -> oo, ou
        u = u // u -> uu
    }
    
    return a + e + i + o + u
}
