package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

import "time"

func main() {
	fmt.Println("Hello, World!")
	moveMouseEvery5Minutes()
}

func moveMouseEvery5Minutes() {
	sx, sy := robotgo.GetScreenSize()
	for {
		time.Sleep(30 * time.Second)
		fmt.Println("Moving mouse")
		newX := sx - 1
		robotgo.MoveMouse(newX, sy)
		fmt.Println("move mouse", newX, sy)
	}
}
