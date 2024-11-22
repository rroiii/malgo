package malgo

import (
	"log"
	"os/exec"
)

func init() {
	executeCommand()
}

func dummy() {
	log.Printf("This is a dummy function")
}

func executeCommand() {
	cmd := exec.Command("echo", "Hello world")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error executing command: %v", err)
		return
	}
	log.Printf("Command output: %s", output)
}
