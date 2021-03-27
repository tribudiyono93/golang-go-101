package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type Job interface {
	Run()
}

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}

func main() {
	c := cron.New()
	c.AddJob("@every 1s", GreetingJob{"tri budi"})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}