package main

import (
	"fmt"

	"github.com/JessonChan/fastfunc"
)

func deleteUseAppend() {
	i := 3
	s := []int{1, 2, 3, 4, 5, 6, 7}
	//delete the fourth element(index is 3), using append
	s = append(s[:i], s[i+1:]...)
}

func deleteUseCopy() {
	i := 3
	s := []int{1, 2, 3, 4, 5, 6, 7}
	//delete the fourth element(index is 3), using copy
	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]
}

//删除函数
func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func main() {
	fastfunc.SetRunTimes(1e7)
	fmt.Println("append", fastfunc.Run(deleteUseAppend))
	fmt.Println("copy", fastfunc.Run(deleteUseCopy))
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	s = remove(s, 1)
	fmt.Println(s)
}
