package term

import (
	"fmt"
	"log"
)

const GREEN = 1
const RED = 2
const YELLOW = 3
const green = "\u001b[32m"
const red = "\u001b[31m"
const yellow = "\u001b[33m"
const reset = "\u001b[0m"

func PCol(msg string, col int) {
	switch col {
	case GREEN:
		fmt.Print(green, msg, reset)
	case RED:
		fmt.Print(red, msg, reset)
	case YELLOW:
		fmt.Print(yellow, msg, reset)
	default:
		fmt.Print(msg)
	}
}

func Ok(msg string) {
	PCol("OK", GREEN)
	if msg != "" {
		fmt.Println(" - ", msg)
	}
}

func Warn(msg string) {
	PCol("WARNING: ", YELLOW)
	fmt.Println(msg)
}

func Error(payload interface{}) {
	PCol("ERROR: ", RED)
	switch payload.(type) {
	case string:
		fmt.Println(payload)
	case Err:
		fmt.Printf("%s. Code: %d", payload.(Err).Message, payload.(Err).Code)
	}
}

func Abort(payload interface{}) {
	Error(payload)
	log.Fatal("Exiting...")
}
