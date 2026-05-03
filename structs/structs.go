package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName  string
	lastName   string
	birthddate string
	age        int
	createdAt  time.Time
	mobile     Mobile // struct defined within the same go file
}

type Mobile struct {
	countryCode string
	number      string
	numberType  string
}

// Here as we are just printing the values not mutating them so recieving them as a pointer dereferece is not mandatory but can be used here as well
func (user User) getUserDetails() {
	fmt.Printf("User Details are : %s %s %s\n", user.firstName, user.lastName, user.birthddate)
}

// Mutuation or Setter methods should take refernces of Reciever arguments as a pointer dereference i.e "*User"
func (user *User) setFirstName(firstName string) {
	user.firstName = firstName
}

func (user User) setFirstNameWithoutMutation(firstName string) {
	user.firstName = firstName
}

func main() {
	//Initialising values to be put in User and Mobile Struct
	userFirstName := "Rohit"
	userLastName := "Walecha"
	userBirthddate := "02-Mar-1994"
	userAge := 32
	userCountryCode := "+91"
	userNumber := "8826057076"
	userNumberType := "Home"
	userMobile := Mobile{
		countryCode: userCountryCode,
		number:      userNumber,
		numberType:  userNumberType,
	}

	var appUser User
	appUser = User{
		firstName:  userFirstName,
		lastName:   userLastName,
		birthddate: userBirthddate,
		age:        userAge,
		createdAt:  time.Now(),
		mobile:     userMobile,
	}

	fmt.Println(appUser)     // Priting the struct as it is
	appUser.getUserDetails() // Priting using getter method which is attached to sturct appUser

	appUser.setFirstNameWithoutMutation("Basant")
	fmt.Println("Setting appUser.firstName to Basant without mutation of actual value not using pointer dereference")
	appUser.getUserDetails() // Priting to check if actual value of appUser.firstName changed or not

	appUser.setFirstName("Basant")
	fmt.Println("Setting appUser.firstName to Basant without mutation of actual value using pointer dereference")
	appUser.getUserDetails() // Priting to check if actual value of appUser.firstName changed or nots
}
