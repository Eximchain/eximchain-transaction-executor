package main

import "os"

func main() {
	switch os.Args[1] {
	case "server":
		RunServerCommand(os.Args[2:])
	case "user":
		RunUserCommand(os.Args[2:])
	}
}
