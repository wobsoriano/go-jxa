package jxa

import (
	"bytes"
	"os/exec"
	"strings"

	mac "github.com/wobsoriano/go-macos-version"
)

func RunJXA(code string) (string, error) {
	mac.AssertMacOSVersion(">=10.10")

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