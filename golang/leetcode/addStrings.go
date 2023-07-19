package main

import "fmt"

func main() {
	fmt.Println(addStrings("13", "129"))
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("11", "11"))
	fmt.Println(addStrings("0", "0"))
	fmt.Println(addStrings("1", "99"))
}

func addStrings(num1 string, num2 string) string {
	ans := ""
	m, n := len(num1)-1, len(num2)-1
	var next byte

	for m >= 0 && n >= 0 {
		sum := (next + num1[m] + num2[n]) % '0'
		next = sum / 10

		ans = string(sum%10+'0') + ans
		m--
		n--
	}

	if m >= 0 {
		for m >= 0 {
			sum := (next + num1[m]) % '0'
			next = sum / 10

			ans = string(sum%10+'0') + ans
			m--
		}
	} else if n >= 0 {
		for n >= 0 {
			sum := (next + num2[n]) % '0'
			next = sum / 10

			ans = string(sum%10+'0') + ans
			n--
		}
	}

	if next == 1 {
		ans = "1" + ans
	}

	return ans
}
