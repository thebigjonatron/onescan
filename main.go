package main

import (
	"onescan/controller"
	"os"
)

func main() {
	args := os.Args[1:]
	controller.Start(args)
}
