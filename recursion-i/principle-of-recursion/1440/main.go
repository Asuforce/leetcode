package main

import "fmt"

func main() {
	reverseString([]byte("hello"))
	reverseString([]byte("Hannah"))
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
	}
	fmt.Println(s)
}
