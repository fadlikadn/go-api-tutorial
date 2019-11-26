package main

import (
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api/utils/email"
)

func main() {
	fmt.Println("Test send email")
	email.SendTestEmail("fadlikadn@hotmail.com")
}
