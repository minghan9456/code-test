// https://leetcode.com/problems/container-with-most-water/submissions/

func maxArea(height []int) int {
    l := 0
    r := len(height) - 1
    max := 0
    
    for l < r {
        
        if height[l] < height[r] {
            if max < height[l] * (r - l) {
               max = height[l] * (r - l)
            }
            l++
        } else {
            if max < height[r] * (r - l) {
               max = height[r] * (r - l)
            }
            r--
        }
    }
    
    return max
}
