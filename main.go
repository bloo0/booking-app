package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	// bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number for tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		// isValidCity := city == "Singapore" || city == "London"
		// isInvalidCity := city != "Singapore" || city != "London"

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			// var noTicketsRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Comeback back next year.")
				break
			}

		} else {
			// fmt.Printf("we only have a total of %v tickets. You can't book %v tickets.\n", remainingTickets, userTickets)
			// fmt.Println("Your input data is invalid, please try again.")
			if !isValidName {
				fmt.Println("first name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("@ sign in email is not present.")
			}
			if !isValidTicketNumber {
				fmt.Printf("we only have a total of %v tickets. You can't book %v tickets.\n", remainingTickets, userTickets)
			}

		}

	}

	//city := "London"
	//switch city {
	//case "New York":
	//	// execute code for booking NY conference tickets
	//case "Singapore":
	//	// execute SG code
	//case "Hongkong":
	//	// exectute HK code
	//case "London", "Berlin":
	//	// exectute code for the said cities
	//default:
	//	fmt.Println("No valid city selected.")
	//}

}
