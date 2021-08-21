// https://leetcode.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
    tmp := map[[26]int][]string{}
    
    for _, str := range strs {
        chars := [26]int{}
        for _, c := range str {
            chars[c - 'a']++
        }
        tmp[chars] = append(tmp[chars], str)
    }
    
    // fmt.Println(tmp)
    
    res := [][]string{}
    for _, list := range tmp {
        res = append(res, list)
    }
    
    return res
}
