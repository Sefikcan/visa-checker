package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		country, city := GetUserInput()
		CheckAppointments(country, city)

		var continueCheck string
		fmt.Print(ContinueCheckMsg)
		_, err := fmt.Scan(&continueCheck)
		if err != nil {
			return
		}
		if strings.ToUpper(continueCheck) != Yes {
			fmt.Println(ExitingMsg)
			os.Exit(0)
		}
	}
}
