package main

import "fmt"

func main(){
	info := map[string]int{
		"rosh": 20,
		"avinash": 21,
		"dorito": 5,
		"bhuvaan": 19,
	}

	for key := range info {
		fmt.Println(key, ":", info[key])
	}

	val, ok := info["dorito"]
	if ok {
		fmt.Println("Dorito is",val, "years old")
	} else {
		fmt.Println("dorito is missing :(")
	}
}