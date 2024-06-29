package main

import (
	"fmt"
)

func divide (a, b float64) (float64, error){
	if(b==0){
		return 0, fmt.Errorf("denominator cannot be 0")
	}
	return a/b, nil
}

func main() {

	ans , _:= divide(10, 9)
	fmt.Println(ans)
	
}