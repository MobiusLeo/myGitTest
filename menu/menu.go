package main

import "fmt"

func main() {
	for {
		var command string
		fmt.Print("MENU：\n1、Welcome\n2、Command\n3、Quit\n")
		fmt.Scan(&command)
		if command == "3" {
			break
		}
		switch command {
		case "1":
			fmt.Println("Hello world!")
		case "2":
			fmt.Println("coming soon…")
		default:
			fmt.Println("Unknown command")
		}
	}
}
