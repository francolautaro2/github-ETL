package main

import (
	"repoInfo/pkg/processing"
	"time"
)

func main() {
	uptimeTicker := time.NewTicker(10 * time.Minute)
	for {
		select {
		case <-uptimeTicker.C:
			processing.Run()
		}
	}

}
