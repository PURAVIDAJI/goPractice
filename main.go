package main

import (
	"fmt"
	"strings"
)

func main(){

	var conferenceName = "Go Conference"
	const conferenceTickets =50 //고정값
	var remainingTickets uint =50
	var bookings =[]string{}

	greetUsers()

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for{

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName :=len(firstName) >=2 && len(lastName) >=2
		isValidEmail := strings.Contains(email, "@")
		isValidEmailTicketNumber := userTickets >0 && userTickets<=remainingTickets

		//isInvalidCity := city !="Singapore" || city !="London"


		if isValidName &&isValidEmail && isValidEmailTicketNumber {
			remainingTickets =remainingTickets -userTickets
		
			bookings = append(bookings, firstName + " " + lastName)
	
			fmt.Printf("Thank you %v %v for booking %v tickets. You will be receiving a confirmation email at %v\n", firstName, lastName, userTickets,email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets,conferenceName)
			firstNames :=[]string{}
			for _ ,booking :=range bookings{
				var names = strings.Fields(booking)
				
				firstNames = append(firstNames,names[0])
	
			}
	
			fmt.Printf("Theses are all our bookings: %v\n" ,firstNames)
			noTicketsRemaining:=remainingTickets==0 
			if noTicketsRemaining{
	
				fmt.Println("Our conference tickets are sold out!")
				break
		
			}

		}else{
			
			if !isValidName{
				fmt.Println("Your name input is wrong")
			}
			if !isValidEmail{
				fmt.Println("Your email is invaild. Include the right format of email.")
			}
			if !isValidEmailTicketNumber{
				fmt.Printf("Your ticket request is invalid, our remaining tickets are %v, adjust to this one.", remainingTickets)
			}

		
		}
		


	}
	

}

func greetUsers () {
	fmt.Println("Welcome to our conference")
}