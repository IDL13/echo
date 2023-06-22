package main

import (
	"os"
	"os/exec"

	"github.com/IDL13/echo/internal/app"
	"github.com/IDL13/echo/pkg/utils"
)

func shellEnv(s *exec.Cmd) {
	s.Stdin = os.Stdin
	s.Stdout = os.Stdout
	s.Stderr = os.Stderr
	s.Run()
}

func StartReactProject() {
	shell := exec.Command("cd", "./../../www/")
	shellEnv(shell)
	shell = exec.Command("npm", "start")
	shellEnv(shell)
}

func main() {
	go StartReactProject()
	a, err := app.New()
	if err != nil {
		utils.Loger(err)
	}

	a.Run()
}
