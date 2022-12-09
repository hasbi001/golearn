package main

import "fmt"

func main(){
	// multiple define constant 
	const (
		country = "Indonesia"
		province = "Jakarta Selatan"
	)

	const counter = 0

	fmt.Println(country);
	fmt.Println(province);	
	fmt.Println(counter);		
}