package helper

func ValidateUserInputs(userTickets uint, firstName string, lastName string, remainingTickets uint) (bool, bool){
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidName := len(firstName) > 2 && len(lastName) > 2
	
	return isValidTicketNumber, isValidName
}