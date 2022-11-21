package helper

import "fmt"

// AdminName is
const AdminName string = "Harjot"

// GreetUser is
func GreetUser(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to Booking App\nGet your virtual tickets for %v from this App, We have total of %v tickets and %v are available\n", conferenceName, conferenceTickets, remainingTickets)
	// *remainingTickets = *remainingTickets * 100
}

// RevokeBooking is
func RevokeBooking() {
	fmt.Printf("User's booking got revoked by ")
}

// InitateMixedSlice is
func InitateMixedSlice() {
	myMixedSlice := []interface{}{1, 2, 3}
	myMixedSlice = append(myMixedSlice, "1")
	myMixedSlice = append(myMixedSlice, "2")
	myMixedSlice = append(myMixedSlice, "3")
	myMixedSlice = append(myMixedSlice, 4)
	myMixedSlice = append(myMixedSlice, 5)
	for _, element := range myMixedSlice {
		fmt.Printf("element is of type %T\n", element)
	}
}
