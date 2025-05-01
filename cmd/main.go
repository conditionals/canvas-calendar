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
var Red = "\033[31m"

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
			if len(entries) == 0 {
			    printRed("Error. No accounts found")
			    break
			}


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
			writeCSV(&entries)
		case 3:
		   printEntries(entries)
		   fmt.Println("What entry would you like to remove? (-1 to exit)")
		   handleRemove(&entries, getInt());

		case 4:
		    printEntries(entries)
		case 5:
		    writeCSV(&entries)
		    os.Exit(0)
		default:
		    printRed("Invalid input.")
		}
	}
}

func promptCLI() int {
	fmt.Println("\nWhat would you like to do?")
	fmt.Println("1. Print due assignments")
	fmt.Println("2. Add a canvas logon")
	fmt.Println("3. Remove a canvas logon")
	fmt.Println("4. List current canvas logons")
	fmt.Println("5. Exit")
	fmt.Print("\nEnter selection: ")

	var num int

	_, err := fmt.Scanln(&num)
	fmt.Println()
	if err != nil || num > 5 || num < 0{
		return -1
	}

	return num
}

func writeCSV(data *[]Entry) {
	file, err := os.Create("list.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := gocsv.MarshalFile(data, file); err != nil {
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
		printRed("Error reading input. Please try again.")
		return getInput()
	}

	return data
}

func getInt() int {
    var data int
    fmt.Printf("Enter your number: ")

    _, err := fmt.Scan(&data)
    if err != nil {
	return getInt()	
    }

    return data
}

func printEntries(entries []Entry) {
	for i := 0; i < len(entries); i++ {
	    fmt.Printf("%d. %s: %s\n", i+1, entries[i].Domain, entries[i].BearerToken)
	}
}

func handleRemove(entries *[]Entry, index int) {
    i := index - 1
    if i < 0 || i >= len(*entries) {
	return 
    }

    *entries = append((*entries)[:i], (*entries)[i+1:]...)
}

func printRed(s string) {
    fmt.Printf("%s%s%s\n", Red, s, Reset)
}
