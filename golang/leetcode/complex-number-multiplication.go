// https://leetcode.com/problems/complex-number-multiplication/

func complexNumberMultiply(num1 string, num2 string) string {
    r1, i1 := divide(num1)
    r2, i2 := divide(num2)
    
    // fmt.Println(r1, i1, r2, i2)
    
    return strconv.Itoa(r1*r2-i1*i2) + "+" + strconv.Itoa(i1*r2 + i2*r1) + "i"
}

func divide(num string) (int, int) {
    idx := strings.Index(num, "+")
    tmpR := num[0:idx]
    tmpI := num[idx+1:len(num)-1]
    
    r, _ := strconv.Atoi(tmpR)
    i, _ := strconv.Atoi(tmpI)
    
    return r, i
}
