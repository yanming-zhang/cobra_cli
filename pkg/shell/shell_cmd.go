package shell

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteShell(cmd []string) (string, error) {
	ret := exec.Command("/bin/bash", cmd...)
	output, err := ret.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func ExecuteShellStdout(cmd []string) error {
	ret := exec.Command("/bin/bash", cmd...)
	ret.Stdout = os.Stdout
	if err := ret.Run(); err != nil {
		fmt.Println("could not run command: ", err)
		return err
	}
	return nil
}
