package internal

import (
	"os/exec"
	"strings"
)

func WriteToClipboardWindows(text string) error {
	cmd := exec.Command("clip")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}
