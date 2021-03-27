package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	c.AddFunc("@hourly", func() {
		fmt.Println("Run hourly")
	})

	c.AddFunc("@daily", func() {
		fmt.Println("Run daily")
	})

	c.AddFunc("@weekly", func() {
		fmt.Println("run weekly")
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
