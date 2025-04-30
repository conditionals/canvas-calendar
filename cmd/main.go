package main

import (
	"os"
	"sort"
	"time"

	"fmt"

	"github.com/conditionals/canvas-calendar/api"
	"github.com/gocarina/gocsv"
)

var Reset = "\033[0m"
var Green = "\033[32m"

func main() {
	tz, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	file, err := os.Open("list.csv")
	if err != nil {
		panic("could not open list.csv")
	}

	defer file.Close()

	var entries []Entry
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		panic(err)
	}

	for true {
		result := promptCLI()
		switch result {
		case 1:
			var assignments []api.CanvasAssignment
			for _, entry := range entries {
				user := &api.CanvasUser{
					Url:   entry.Domain,
					Token: entry.BearerToken,
				}
				assignments = append(assignments, user.FetchAssignments()...)
			}

			printAssignments(assignments, tz)
		case 2:
			var newEntry Entry

			fmt.Println("Please enter the domain of the new institution's canvas.")
			newEntry.Domain = getInput()

			fmt.Println("Please enter the bearer token.")
			newEntry.BearerToken = getInput()

			entries = append(entries, newEntry)
		case 4:
			for i := 0; i < len(entries); i++ {
				fmt.Printf("%s: %s\n", entries[i].Domain, entries[i].BearerToken)
			}
		case 5:
			os.Exit(0)
		}
	}
}

func promptCLI() int {
	fmt.Println("\nWhat would you like to do?")
	fmt.Println("1. Print due assignments (default)")
	fmt.Println("2. Add a canvas logon")
	fmt.Println("3. Remove a canvas logon")
	fmt.Println("4. List current canvas logons")
	fmt.Println("5. Exit")
	fmt.Print("Enter selection: ")

	var num int

	_, err := fmt.Scanln(&num)
	fmt.Println()
	if err != nil {
		return 1
	}

	if num > 5 || num < 0 {
		return 1
	}

	return num
}

func writeCSV(data []*Entry) {
	file, err := os.Create("list.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := gocsv.MarshalFile(&data, file); err != nil {
		panic(err)
	}

}

type Entry struct {
	Domain      string `csv:"domain"`
	BearerToken string `csv:"bearer_token"`
}

func printAssignments(assignments []api.CanvasAssignment, tz *time.Location) {
	sort.Slice(assignments, func(i, j int) bool {
		return assignments[i].Assignment.DueAt.Before(assignments[j].Assignment.DueAt)
	})

	fmt.Println("Due Assignments: -------------------------")
	for _, x := range assignments {
		out := fmt.Sprintf("Class: %s | Due: %s (%s), Name: %s", api.GetCourseName(x.Assignment.CourseID), x.Assignment.DueAt.In(tz).Format("Jan 2 03:04 PM"), x.Assignment.DueAt.In(tz).Format("01/02/06"), x.Assignment.Name)

		if x.Assignment.UserSubmitted || x.Assignment.GradedSubmissionsExist {
			out = fmt.Sprintf("%s%s%s", Green, out, Reset)
		}

		fmt.Println(out)
	}
}

func getInput() string {
	var data string
	_, err := fmt.Scanln(&data)
	if err != nil {
		fmt.Println("Error reading input. Please try again.")
		return getInput()
	}

	return data
}
