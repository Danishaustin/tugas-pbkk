package main

import "fmt"

func main() {
	var input int

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&input)

	if input == 42 {
		fmt.Println("Hello Universe")
	} else {
		fmt.Println(input)
	}

}
