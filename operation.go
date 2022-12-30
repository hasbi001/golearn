package main

import "fmt"

func main(){
	// 1. Matematika + - / % 
	var no1 int32;

	no1 = 10;
	no1 += 4; 
	fmt.Println(no1);

	var no2 int32;

	no2 = no1 - 3;
	fmt.Println(no2);
	no2++
	fmt.Println(no2%2);

	// 2. Perbandingan < > == != >= <=
	if no1 != no2 {
		fmt.Println("true");	
	} else {
		fmt.Println("false");
	}
}