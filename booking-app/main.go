package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remaningTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]userData, 0)

var wg = sync.WaitGroup{}

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greatUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickedNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remaningTickets)

		if isValidName && isValidEmail && isValidTickedNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstName()
			fmt.Printf("first name are %v\n", firstNames)

			if remaningTickets == 0 {
				fmt.Println("Our conference is booked out.Come back next year")
				break
			}
		} else {
			fmt.Printf("Your Input data is invalid\n")
			continue
		}

	}
}

func greatUser() {
	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Printf("We have Total of %v tickets and %v are still remaning\n", conferenceTickets, remaningTickets)
	fmt.Println("Get your Tickets here to attend Conference")
}

func getFirstName() []string {
	firstNames := []string{}

	// getting data from slice
	// for _, booking := range bookings {
	// 	firstNames = append(firstNames, strings.Fields(booking)[0])
	// }

	//getting data from map
	// for _, booking := range bookings {
	// 	firstNames = append(firstNames, booking["firstName"])
	// }

	// geetting data from struct
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	println("Enter your First Name:\n")
	fmt.Scan(&firstName)
	println("Enter your last Name:\n")
	fmt.Scan(&lastName)
	println("Enter your email:\n")
	fmt.Scan(&email)
	println("Enter Number of tickets you want to buy:\n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remaningTickets = remaningTickets - userTickets

	// example to creeate map.limitation for map is it can only have one data type.
	// userData := make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//example with struct
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("The booking details are as follow %v\n", bookings)
	fmt.Printf("Thankyou %v %v for booking %v tickets.You will recive conformation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaning for %v\n", remaningTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v \nto email address%v\n", userTickets, firstName, lastName, email)
	fmt.Printf("Sending tickets:\n%v\nto email address %v\n", tickets, email)
}
