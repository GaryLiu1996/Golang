package main

import "fmt"

//在Go中按值传递时，为什么有时会更改切片？

func main() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	fmt.Println(cap(s))
	reverse(s)
	fmt.Println(s)
}
func reverse(s []int) {
	s = append(s, 999)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
