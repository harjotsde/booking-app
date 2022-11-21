package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"booking-app/helper"
)

const conferenceTickets uint = 50

var (
	remainingTickets uint = conferenceTickets
	// bookings         []map[string]string
	bookings       []UserData
	conferenceName string = "Go Conference"
	wg             sync.WaitGroup
)

// UserData is
type UserData struct {
	userName        string
	email           string
	numberOfTickets uint
}

func main() {
	// bookings := []string{"jack"}
	// fmt.Printf("Conference tickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to Booking App\nGet your virtual tickets for", conferenceName, "from this App, We have total of", conferenceTickets, "tickets and", remainingTickets, "are available")
	helper.GreetUser(conferenceName, conferenceTickets, remainingTickets)
	startBooking()
	helper.RevokeBooking()
	fmt.Printf("the order of %v\n", helper.AdminName)

	// An example for mixed type slice
	// helper.InitateMixedSlice()
}

func getUserInputs() (string, string, uint) {
	var userName, email string
	var userTickets uint
	fmt.Printf("Enter your Name\n")
	fmt.Scan(&userName)
	fmt.Printf("Enter your email\n")
	fmt.Scan(&email)
	fmt.Printf("Enter number of tickets you need\n")
	fmt.Scan(&userTickets)
	return userName, email, userTickets
}

func isValidBooking(name, email string) bool {
	return len(name) > 2 && strings.Contains(email, "@")
}

func startBooking() {
	for {
		userName, email, userTickets := getUserInputs()

		if !isValidBooking(userName, email) {
			fmt.Printf("Booking validation failed, Provided details are Invalid\n")
			break
		}
		// userData := make(map[string]string)
		// userData["userName"] = userName
		// userData["email"] = email
		// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
		userData := UserData{
			userName:        userName,
			email:           email,
			numberOfTickets: userTickets,
		}
		mixedData := make(map[string]interface{})
		mixedData["userName"] = userName
		mixedData["email"] = email
		mixedData["userTickets"] = userTickets
		bookings = append(bookings, userData)
		// fmt.Println(mixedData)
		// fmt.Println(userData)

		if userTickets > remainingTickets {
			userTickets = remainingTickets
		}

		fmt.Printf("%v booked %v tickets\nThankyou for booking %v tickets, you will soon receive a confirmation email at %v\n", userData.userName, userTickets, userTickets, email)
		wg.Add(1)
		go sendTicket(userData)
		remainingTickets = remainingTickets - userTickets

		// manipulatedBookings := []string{}

		// for index, bookedBy := range bookings {
		// 	manipulatedBookings = append(manipulatedBookings, bookedBy+" "+strconv.Itoa(index))
		// }

		// fmt.Println(manipulatedBookings)
		fmt.Println(bookings)
		// bookings[0] = userName

		fmt.Printf("%v have %v tickets remaining\n", conferenceName, remainingTickets)

		if remainingTickets == 0 {
			fmt.Printf("%v tickets are sold out for this season\n", conferenceName)
			break
		}
	}
	wg.Wait()
}

func sendTicket(userData UserData) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets, reserved for %v", userData.numberOfTickets, userData.userName)
	fmt.Println("################")
	fmt.Printf("%v are sent at %v\n", ticket, userData.email)
	fmt.Println("################")
	wg.Done()
}
