// https://leetcode.com/problems/sign-of-the-product-of-an-array/

func arraySign(nums []int) int {
    sign := 1
    for i := 0; i < len(nums); i++{
        if nums[i] < 0 {
            sign = -sign
        }
        
        if nums[i] == 0 {
            return 0
        }
    }
    
    return sign
}
