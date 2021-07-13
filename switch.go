package main

import (
	"fmt"
	"time"
)

func main() {

	i := 4
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's sun light")
	default:
		fmt.Println("It's noon light")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I'm a int")
		case bool:
			fmt.Println("I'm an blool")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(1)
	whatAmI(true)
	whatAmI("hey")
}
