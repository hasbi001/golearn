package main

import (
	"fmt"
	"strconv"
)

func main(){
	var (
		v32 int32 = 129
		v8 int8 = int8(v32)
		v64 int64 = int64(v32)
	)
	int := 334 
	fmt.Println(v32);
	fmt.Println(v8);
	fmt.Println(v64);

	str := strconv.Itoa(int);
	fmt.Println(str)
}