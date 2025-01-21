package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client *http.Client
var courseMap map[int]string

func init() {
    client = &http.Client{}
    courseMap = make(map[int]string)
}


func (u *CanvasUser) FetchAssignments() []CanvasAssignment {
    req, err := http.NewRequest("GET", u.formatCalendarRequestURL(), nil)
    if err != nil {
	panic("failed to create request")
    }

    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", u.Token))

    resp, err := client.Do(req)
    if err != nil {
	panic(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
	fmt.Printf("warn: status code %d\n", resp.StatusCode)
    }
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
	panic(err)
    }
    
    var assignmentList []CanvasAssignment

    err = json.Unmarshal(body, &assignmentList)
    if err != nil {
	panic(err)
    }

    return assignmentList
}


// returns a psuedo-url-encoded list of class codes for easy formatting
func (u *CanvasUser) fetchFormattedClasses() string{
    classCodes := ""

    req, err := http.NewRequest("GET", u.formatClassRequestURL(), nil)
    if err != nil {
	panic(err)
    }

    req.Header.Add("Authorization", "Bearer " + u.Token)

    resp, err := client.Do(req)
    if err != nil {
	panic(err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
	panic(err)
    }

    var classList classes

    err = json.Unmarshal(body, &classList)
    if err != nil {
	panic(err)
    }

    for _, x := range classList {
	classCodes += fmt.Sprintf("&context_codes[]=course_%d", x.ID)
	courseMap[x.ID] = x.Name 
    }

    return classCodes
}

func (u *CanvasUser) formatCalendarRequestURL() string {
    endpointURL := fmt.Sprintf("%s/api/v1/calendar_events", u.Url) 

    endDate := time.Now().Add(time.Hour * 24 * 7).Format(time.RFC3339)

   withDates := fmt.Sprintf("%s?end_date=%s&type=assignment%s", endpointURL, endDate, u.fetchFormattedClasses())

    return withDates 
}

func (u *CanvasUser) formatClassRequestURL() string {
    return fmt.Sprintf("%s/api/v1/courses", u.Url)
}

func GetCourseName(id int) string {
    return courseMap[id]
}
