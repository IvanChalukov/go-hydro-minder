package main

import (
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	for {
		time.Sleep(60 * time.Minute)

		err := beeep.Notify("Water Reminder", "Time to dring water!", "")
		if err != nil {
			panic(err)
		}
	}
}
