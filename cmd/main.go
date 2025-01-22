package main

import (
	"os"
	"sort"
	"time"

	"fmt"

	"github.com/conditionals/canvas-calendar/api"
	"github.com/joho/godotenv"
)

var Reset = "\033[0m" 
var Green = "\033[32m" 

func main() {
    err := godotenv.Load()
    if err != nil {
	panic(err) 
    }

    tz, err := time.LoadLocation(os.Getenv("TIMEZONE"))
    if err != nil {
	panic(err)
    }


    sf_token := os.Getenv("SF_BEARER_TOKEN")
    sf_url := os.Getenv("SF_CANVAS_URL")

    uf_token := os.Getenv("UF_BEARER_TOKEN")
    uf_url := os.Getenv("UF_CANVAS_URL")

    sf := &api.CanvasUser{
	Url: sf_url,
	Token: sf_token,
    }

    uf := &api.CanvasUser{
        Url: uf_url,
        Token: uf_token,
    }
    
    assignments := append(sf.FetchAssignments(), uf.FetchAssignments()...)

    // sort based on due date to return earliest first
    sort.Slice(assignments, func(i, j int) bool {
        return assignments[i].Assignment.DueAt.Before(assignments[j].Assignment.DueAt)
    })

    fmt.Println("Due Assignments: -------------------------")
    for _, x := range assignments {
	out := fmt.Sprintf("Class: %s | Due: %s (%s), Name: %s", api.GetCourseName(x.Assignment.CourseID),  x.Assignment.DueAt.In(tz).Format("Jan 2 03:04 PM"), x.Assignment.DueAt.In(tz).Format("01/02/06"), x.Assignment.Name)
	
	if x.Assignment.UserSubmitted || x.Assignment.GradedSubmissionsExist {
	    out = fmt.Sprintf("%s%s%s", Green, out, Reset)
	}

	fmt.Println(out)
    }

}
