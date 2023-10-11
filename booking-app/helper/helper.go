package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remaningTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickedNumber := userTickets > 0 && userTickets <= remaningTickets

	return isValidName, isValidEmail, isValidTickedNumber
}
