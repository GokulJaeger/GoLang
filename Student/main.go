package main

import (
	"flag"
	"fmt"
	"strconv"
)

var student [3][11]string

// "Roll Number", "Name", "Subject 1", "Subject 2", "Subject 3", "Subject 4", "Subject 5", "Total", "Average", "Number of failed Subjects", "Grade", "Failed marks"
var quest = [7]string{"Roll Number", "Name", "Subject 1", "Subject 2", "Subject 3", "Subject 4", "Subject 5"}
var rw = 0

func addstd() string {
	fmt.Println("Enter 3 Students details!.")
	for i := 0; i < 3; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf(quest[j])
			fmt.Scanln(&student[i][j])
		}
	}
	return "Add Function"
}

func viewstd() string {

	if student[0][0] == "" {
		fmt.Println("No records found!")
	} else {
		fmt.Println("Student List")
		for i := 0; i < 3; i++ {
			for j := 0; j < 11; j++ {
				fmt.Println(student[i][j])
			}
		}
	}
	return "View Student Function"
}

func searchstd(a int) string {
	flag.Parse()

	for i := 0; i < 3; i++ {
		for j := 0; j < 11; j++ {
			v, _ := strconv.Atoi(student[i][j])
			if v == a {

			}
		}
	}
	return "Search Student Function"
}

func main() {

	var choice int
	var roll int
	for {
		fmt.Println("1.Add Student  2.View Students  3.Search  4.Exit")
		fmt.Println("Enter your choice")
		fmt.Scanln(&choice)

		if choice == 1 {
			st1 := addstd()
			fmt.Printf(st1)
		} else if choice == 2 {
			st2 := viewstd()
			fmt.Printf(st2)
		} else if choice == 3 {
			fmt.Println("Enter the roll number")
			fmt.Scanln(&roll)
			st3 := searchstd(roll)
			fmt.Printf(st3)
		} else {
			fmt.Println("Thanks for using student portal.")
		}
	}

}
