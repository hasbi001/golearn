package main

import "fmt"

func main(){
	var monts = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	slice1 := monts[4:7]
	fmt.Println(monts)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0]="mei syah"
	fmt.Println(monts)

	append1 := append(slice1, "desember ceria")
	fmt.Println(append1)
}