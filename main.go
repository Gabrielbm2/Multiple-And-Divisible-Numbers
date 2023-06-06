package main

import "fmt"

func main() {
	fmt.Println("Parte 1")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Parte 2")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
