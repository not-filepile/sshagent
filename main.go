package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	host := ""

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-a1":
			host = ""
		case "-s1":
			host = ""
		default:
			log.Fatalf("Invalid argument: %s. Use -a1 or -s1.", os.Args[1])
		}
	}

	user := "ubuntu"
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", user, host))

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute SSH command: %s", err)
	}
}
