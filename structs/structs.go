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

	fmt.Println(appUser)
}
