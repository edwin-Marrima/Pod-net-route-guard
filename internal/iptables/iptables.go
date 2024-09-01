package iptables

import (
	"bytes"
	"io"
	"os/exec"
)

const iptable = "iptable"

type CmdError struct {
	cmd      exec.Cmd
	msg      string
	exitCode int
}

func (e *CmdError) Error() string {
	return ""
}

type IPTables struct {
}

func (ipt *IPTables) runWithOutput(args []string, stout io.Writer) error {
	args = append([]string{iptable}, args...)

	var stderr bytes.Buffer

	cmd := exec.Cmd{
		Path:   iptable,
		Args:   args,
		Stdout: stout,
		Stderr: &stderr,
	}
	if err := cmd.Run(); err != nil {
		switch e := err.(type) {
		case *exec.ExitError:
			return &CmdError{cmd: cmd, msg: e.String(), exitCode: 1}
		default:
			return err
		}

	}
	return nil
}

func (ipt *IPTables) ListChains(table string) ([]string, error) {
	//cmd := []string{"-t", "table", "-S"}

	return nil, nil
}
