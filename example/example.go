package main

import (
	"common"
	"time"
)

func main() {
	for {
		common.Debug("test debug")
		common.Info("test info")

		time.Sleep(time.Minute)
	}
}
