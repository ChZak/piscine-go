package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

type Door struct {
	state bool
}

// si cest ouvert true
// si c fermer false

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = true
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	ptrDoor.state = false
	return false
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	ptrDoor.state = false
	return false
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Open Door\n")
	ptrDoor.state = true
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if !IsDoorClose(door) {
		OpenDoor(door)
	}
	if !IsDoorOpen(door) {
		CloseDoor(door)
	}
}
