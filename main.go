package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	mins := flag.Int("mins", 60, "awaiting minutes (int)")
	flag.Parse()

	for {
		fmt.Printf("Next notification will be after %d minutes.", *mins)
		time.Sleep(time.Duration(*mins) * time.Minute)

		err := beeep.Notify("Water Reminder", "Time to dring water!", "")
		if err != nil {
			panic(err)
		}
	}
}
