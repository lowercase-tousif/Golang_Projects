package handler

import (
	"fmt"
	"student-management/model"

	"github.com/fatih/color"
)

func AddStudent() {
	var s model.Student

	fmt.Println("Enter the student name: ")
	fmt.Scan(&s.Name)

	fmt.Println("Enter the student age: ")
	fmt.Scan(&s.Age)

	fmt.Println("Enter the student gender: ")
	fmt.Scan(&s.Gender)

	fmt.Println("Enter 3 hobbies:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Hobby %d: ", i+1)
		fmt.Scan(&s.Hobbies[i])
	}

	model.Students = append(model.Students, s)
	fmt.Print("\n\t\t✅ Student Added Successfully\n\t\t\n")
}

func ViewStudents() {
	if len(model.Students) == 0 {
		color.Red("\n🚫 No students found.\n")
		return
	}

	for i, s := range model.Students {
		fmt.Println()
		color.Cyan("🎓 Student %d Details", i+1)
		color.Yellow("═════════════════════════════════════")
		color.Green("📛 Name     : %s", s.Name)
		color.Green("🎂 Age      : %d", s.Age)
		color.Green("🚻 Gender   : %s", s.Gender)
		color.Green("🎯 Hobbies  :")
		for _, h := range s.Hobbies {
			fmt.Printf("   ➡️  %s\n", h)
		}
		color.Yellow("═════════════════════════════════════")
	}
}

func Menu() {
	for {
		fmt.Println("========= Menu =========")
		fmt.Println("1. Add Student")
		fmt.Println("2. Show All Students")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			AddStudent()
		case 2:
			ViewStudents()
		case 5:
			fmt.Println("👋 Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
