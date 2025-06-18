package main

import (
	"errors"
	"fmt"
)

func sum(n int) (int, error){

	if n < 0{
		return 0, errors.New("input less than zero")
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum, nil
}

func main(){
	var n int
	fmt.Scanln(&n)
	res, err := sum(n)
	if err != nil {
		fmt.Println("error", err)
	}
	if res > 50{
		fmt.Println("rosh")
	}else {
		fmt.Println("HEJOSKDJ")
	}
}