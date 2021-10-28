package fzf

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/drewstinnett/azurectx-go/internal/subscription"
)

func withFilter(command string, input func(in io.WriteCloser)) []string {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}
	cmd := exec.Command(shell, "-c", command)
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	go func() {
		input(in)
		in.Close()
	}()
	result, _ := cmd.Output()
	return strings.Split(string(result), "\n")
}

func PickSubscription() (string, error) {
	subs, err := subscription.GetSubscriptionNames()
	if err != nil {
		return "", err
	}
	filtered := withFilter("fzf", func(in io.WriteCloser) {
		for _, p := range subs {
			fmt.Fprintln(in, p)
		}
	})
	return filtered[0], nil
}
