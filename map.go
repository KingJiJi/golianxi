package main

import "fmt"
import "strings"

func mapchanged(h map[string]int, k string, v int) {
	h[k] = v
}

func listchanged(h []string, s string) {
	h = append(h, s)
}
func listchanged1(h *[]string, s string) {
	*h = append(*h, s)
}
func main() {
	hahamap := make(map[int][]int)

	hahamap[0] = append(hahamap[0], 90)
	hahamap[0] = append(hahamap[0], 900)
	hahamap[0] = append(hahamap[0], 9000)
	fmt.Printf("%#v\n", hahamap)

	halist := hahamap[0]
	halist = append(halist, 888)
	halist = append(halist, 88)
	halist = append(halist, 8)
	fmt.Printf("%#v\n", hahamap)
	hahamap[1] = append(halist, 77)
	halist = append(halist, 77)
	hahamap[3] = append(halist, 54)
	hahamap[2] = append(halist, 6)
	fmt.Printf("%#v\n", hahamap)
	for k, v := range hahamap {
		fmt.Printf("key:%d, velue:%#v\n", k, v)
	}

	mymm := make(map[string]int)
	mapchanged(mymm, "haha", 0)
	mapchanged(mymm, "haha1", 33)
	mapchanged(mymm, "haha2", 50)
	mapchanged(mymm, "haha3", 780)
	for k, v := range mymm {
		fmt.Printf("mymm key:%s, velue:%#v\n", k, v)
	}

	hlist := []string{"hello"}
	listchanged(hlist, "world")
	listchanged(hlist, "......")

	hlistall := strings.Join(hlist, " ")
	fmt.Printf("hlistall %s \n", hlistall)
	listchanged1(&hlist, "world")
	listchanged1(&hlist, "......")

	hlistall = strings.Join(hlist, " ")
	fmt.Printf("hlistall %s \n", hlistall)

	return

}
