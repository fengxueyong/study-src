package main // it's must for main go

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() { // it's a must for main function
	fmt.Println("Hello,World!")
	names := []string{"fengxueyong", "zhangsan", "lisi"}
	name_maps, err := greetings.Hellos(names)

	if err != nil {
		// will simul..stop the program
		log.Fatal("hello fatal")
	}

	fmt.Println(name_maps)
}
