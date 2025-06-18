package main

import (
	"fmt"
	"sort"
)

func mapp(){
	varia := map[int]string{
		1 : "rosh",
		2 : "gayboi",
		3 : "hii",
	}
	name := varia[1]
	fmt.Println(name)
}

func main(){
	fmt.Print("hiii\n")
	new := []int {1,7,8,2,4};
	sort.Ints(new)
	mapp()
}