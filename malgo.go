package malgo

import (
	"log"
	"os"
	"os/exec"
)

func run() {
	cmd := exec.Command("echo", "hello world")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
