package main

import (
	"fmt"
	"github.com/bamzi/jobrunner"
)

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 5s", ReminderEmails{})
}

type ReminderEmails struct {

}

func (e ReminderEmails) Run() {
	fmt.Printf("Every 5 sec send reminder emails \n")
}
