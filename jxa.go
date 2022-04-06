package jxa

import (
	"bytes"
	"os/exec"
	"strings"
)

func RunJXA(code string) (string, error) {
	cmd := exec.Command("osascript", "-l", "JavaScript")
	cmd.Stdin = strings.NewReader(code)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	
	data := strings.TrimSpace(out.String())
	return data, nil
}