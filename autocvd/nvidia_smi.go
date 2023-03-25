package autocvd

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var execCommand = exec.Command

func ShellCommand(name string, arg ...string) (string, error) {
	cmd := execCommand(name, arg...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}

func GetInstalledGpus() (int, error) {
	out, err := ShellCommand("/usr/bin/env", "nvidia-smi", "-L")
	if err != nil {
		return 0, err
	}
	return len(strings.Split(out, "\n")), nil
}

func GetIsFree(gpu int) (bool, error) {
	out, err := ShellCommand(
		"/usr/bin/env", "nvidia-smi",
		"--query-compute-apps=pid",
		"--format=csv,noheader,nounits",
		"-i", fmt.Sprint(gpu),
	)
	if err != nil {
		return false, err
	}
	return out == "", nil
}

func GetFreeGpuMemory(gpu int) (int, error) {
	out, err := ShellCommand(
		"/usr/bin/env", "nvidia-smi",
		"--query-gpu=memory.free",
		"--format=csv,noheader,nounits",
		"-i", fmt.Sprint(gpu),
	)
	if err != nil {
		return 0, err
	}

	mem, err := strconv.Atoi(out)
	if err != nil {
		return 0, err
	}
	return mem, nil
}
