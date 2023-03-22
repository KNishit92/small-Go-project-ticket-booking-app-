// all the code must be in a package,  the first line in Go file must be "package...". Go provides various core packages for utility.
// a package is a collection of source files
package main

// imports
import (
	"fmt"
	"sync" // ********* provides basic synchronization functionality
	// "strconv"
	// "strings"
	"booking-app/helper"
	"time"
)

// ************** //
// package level variables defined outside all functions. accessible to all functions and files in the same package
const conferenceTickets = 50
var conferenceName string = "Go Conference" // explicit type decalaration
var RemainingTickets uint = 50 // so that it can't take -ve values
// var bookings []string // slice
// var bookings =  make([]map[string]string, 0) // empty slice of maps, 0 is the initial size of the list of maps
var bookings =  make([]UserData, 0)

// struct data types can hold mixed data types, it's like a lightweight Class that eg. doesn't support inheritance
// type keyword creates a new type with specified name
// ex: create a type UserData using the struct firstName, lastName ...
type UserData struct { 
	firstName string
	lastName string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// give a starting/entry pt to the Go compiler. the "main" fuction is the entry pt of a Go program
func main() {
	
	// Println to print on a new line
	// fmt.Print("hello") 
	// fmt.Print("world")
	// fmt.Println("*******") 
	// fmt.Println("hello") 
	// fmt.Println("world")

	// Println prints new line automatically
	// fmt.Println("Welcome to our", conferenceName  , "booking app") 
	// fmt.Println("We have a total of ", conferenceTickets  , " tickets and ", remainingTickets, " are available.") 
	// fmt.Println("Get your tickets here to attend")

	greetUser()

	// print types
	fmt.Printf("conferenceTickets type is %T and conferenceName type is %T.\n", conferenceTickets, conferenceName) 

	// ******************* Wait Group waits for the launched go routine to complete execution ****************


	
	// store user names in an array of strings of size 50
	// ---------------  var bookings [50]string -----------------------
	// slices are useful when length of array is not known, since it's dynamic in nature, they are also indexed
	// like arrays, and have a size which gets resized when required
	// for {
		
		firstName, lastName, userTickets := userInput()
		isValidTicketNumber, isValidName := helper.ValidateUserInputs(userTickets, firstName, lastName, RemainingTickets)

		if isValidTicketNumber && isValidName {

		bookTickets(firstName,lastName,userTickets)
		// make this concurrent using "go" keyword. 
		// ********* Main go routine doesn't wait for any other thread **********

		// use this before creatiung a new thread to make the main thread wait for the execution of the other go rountines.
		wg.Add(1) // ************* sets the number of go routines to wait for , increases the wait grp counter by 
				  // the amt passed as param
		go sendTicket(userTickets, firstName, lastName)
		
		firstNames := getFirstNames()
		// underscores are used to indicate unused variables
		
		
		fmt.Printf("These are the names of the users who booked: %v\n", firstNames)
		fmt.Println(&RemainingTickets)
		if RemainingTickets == 0 {
			println("Tickets sold out. Come back  next year!!")
			// break
		} 

	} else {
		if !isValidTicketNumber {
			fmt.Println("Entered no of ticketd is invalid, try again !")
		}
		if !isValidName {
			fmt.Println("Entered name is invalid, try again !")
		}
	}	
// }
	wg.Wait() // ****************** waits till the wait group counter is 0 ****************
}
// pointer is a variable that points to the memory address of another variable


func greetUser() { // access repeated params from package level variables, for cleaner code
	fmt.Printf("Welcome to our conference: %v\n", conferenceName)
	// using placeholder \n is for new line
	// fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are available.\n", conferenceTickets, RemainingTickets) 
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	var firstNames []string
	for _, booking := range(bookings) {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func userInput() (string, string, uint) {
	var firstName string // (since we are not immediately assigning any value to the variable, hence data type needs to be mentioned)
	var userTickets uint
	var lastName string
	// var city string
	// get user passed value and assign to userName
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	// instead of hard coding as below, ask for user name as above
	// userName = "Tim"
	// userTickets = 2
	fmt.Println("Enter no of tickets you want to book: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint) {
	// create maps for user data
	// var userData = make(map[string]string)  // with [] represents keys type *********** can't have mixed data types for keys
	// create struct type UserData
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// // convert int to string
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // decimal pt to base 10
	// add elements to slice
	bookings = append(bookings, userData)

	fmt.Println(bookings)

	fmt.Printf("Thank you %v for booking %v tickets. Your tickets are confirmed.\n", firstName + " " + lastName, userTickets)
	// pointer returns the memory location of remainingTickets
	RemainingTickets -= userTickets
	fmt.Printf("No of tickets remaining are: %v for %v\n", RemainingTickets, conferenceName)
}

func sendTicket (userTickets uint, firstName string, lastName string) {
	// simulate a delay of 10 secs for sending email to a user
	time.Sleep(10 * time.Second) // hold on the execution of below lines for 10 secs
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("*********************")
	fmt.Printf("Sending tickets to: %v at registered email address\n", tickets)

	wg.Done() // ********* decreases wait grp counter by 1 ********* , go routine calls to indicate it's finished
	 		  //  and thus the main func/thread doesn't have to wait for it anymore
}