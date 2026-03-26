package main

import "fmt"

func panicing() {
	panic("Ou shit!...")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoveed. Error: \n", r)

		}
	}()

	panicing()
	fmt.Println("Check")

}
