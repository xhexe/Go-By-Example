package main

import "fmt"

func main() {
	i := 0
	for {
		i ++

		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 20 {
			break
		}
	}
	fmt.Println([]byte("Dupa"))
}
