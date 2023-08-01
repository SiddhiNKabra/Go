package main

import "fmt"

func main() {
	conf := "Go Conference"
	const tickets = 50
	var remainingtickets uint = 50
	if remainingtickets == 0 {
		fmt.Println("REGRET\n!!No more bookings!!\nAll seats are occupied")

	} else {
		fmt.Printf("Welcome to %v booking application.\n", conf)
		fmt.Printf("We have a total of %v tickets and %v tickets are remaining.\n", tickets, remainingtickets)
		fmt.Println("Get your tickets")

		var firstname string
		var lastname string
		var mailid string
		var usertickets uint
		fmt.Print("Enter number of tickets you want to book: ")
		fmt.Scan(&usertickets)

		remainingtickets = remainingtickets - usertickets

		x := 0
		for x < int(usertickets) {
			fmt.Print("Enter your first name: ")
			fmt.Scan(&firstname)
			fmt.Print("Enter your last name: ")
			fmt.Scan(&lastname)
			fmt.Print("Enter your mail-id: ")
			fmt.Scan(&mailid)
			fmt.Printf("Your booking has been confirmed\n")
			x++
		}
		fmt.Printf("%v tickets are remaining for %v \n", remainingtickets, conf)

	}

}
