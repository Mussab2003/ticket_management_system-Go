//This package tells the starting point of the app. 
package main

//contains I/O commands
import "fmt"


//Main Function similar to one in C / C++
func main() {
	//Different Ways to create a variable in Go
	var conferenceName = "Go Conference" //The classic way
	const totalTickets = 50              //To declare a variable whose value cannot be changed
	remainingTickets := totalTickets     //It is equal to var remaingTickets = totalTickets (an easy and efficient way to declare var variables)
	//Go is static typed language so we have to declare the datatype at compile time.
	//So if we initiaze and declare variable at the sametime then it automatically infers the datatype
	//else we have to explicitly tell the datatype.
	var firstName string                 
	var lastName string
	var emailAddress string
	var userTickets int

	//To get the type of a variable use the %T in the printf statement

	//Different type of print functions
	fmt.Println("Welcome to ", conferenceName, " booking application")
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("Total Tickets: %v, Remaining Tickets: %v\n", totalTickets, remainingTickets)
	
	fmt.Printf("Enter First Name: ")
	//For taking the input
	fmt.Scan(&firstName)
	fmt.Printf("Enter Last Name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter Email Address: ")
	fmt.Scan(&emailAddress)
	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("ThankYou for booking. Your information are: \n First Name: %v\n Last Name: %v\n Email Address: %v\n Number of Tickets: %v\n", firstName, lastName, emailAddress, userTickets)
	remainingTickets -= userTickets

	//Arrays in Go
	var bookings[50] string

	bookings[0] = "Mussab"

	//Slices are a sort of array in which we don't give the size. So it can automatically increase the size. 
	var bookings2[] string
	//To add value in a slice
	bookings2 = append(bookings2, "Mussab")
	
}
