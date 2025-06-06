package main

import "fmt"

func main() {
	productsInMemory, err := LoadData()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(productsInMemory)

}
