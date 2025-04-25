package userinput

import "fmt"

func Input() {
	var userName string
	var userAge int

	fmt.Scan(&userName)
	fmt.Scan(&userAge)
}
