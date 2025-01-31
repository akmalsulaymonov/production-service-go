package main

import "fmt"

// Rin - is going to be responsible for
// the instantiation and startup fot our
// go application
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REST API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
