package main

import (
	"fmt"
)

func main() {

	var s []byte

	s = []byte("ABCD")
	reverseString(s)
	fmt.Println(string(s))

}

func reverseString(s []byte) {

	right := len(s) - 1

	for left := 0; left < len(s)/2; left++ {
		s[left], s[right] = s[right], s[left]
		right -= 1
	}
}

func reverseString(s []byte)  {
    l := len(s) - 1
    m := l/2
    if l >= 1 {
        for i := 0; i <= m; i++ {
            s[i], s[l-i] = s[l-i], s[i]
        }
    }
}

/*
func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--

	}
}
*/
