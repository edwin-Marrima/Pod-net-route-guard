package iptables

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

const iptable = "iptables"

type CmdError struct {
	cmd      exec.Cmd
	msg      string
	exitCode int
}

func (e *CmdError) Error() string {

	return fmt.Sprintf("running %v: exit status %v: %v", e.cmd.Args, e.exitCode, e.msg)

}

type IPTables struct {
	path string
}

func New() (*IPTables, error) {
	path, err := exec.LookPath(iptable)
	if err != nil {
		return nil, err
	}
	return &IPTables{
		path: path,
	}, nil
}

func (ipt *IPTables) runWithOutput(args []string, stout io.Writer) error {
	args = append([]string{iptable}, args...)

	var stderr bytes.Buffer

	cmd := exec.Cmd{
		Path:   ipt.path,
		Args:   args,
		Stdout: stout,
		Stderr: &stderr,
	}
	if err := cmd.Run(); err != nil {
		var e *exec.ExitError
		switch {
		case errors.As(err, &e):
			return &CmdError{cmd: cmd, msg: stderr.String(), exitCode: 1}
		default:
			return err
		}

	}
	return nil
}

func (ipt *IPTables) ListChains(table string) ([]string, error) {
	cmd := []string{"-t", table, "-S"}
	var stdout bytes.Buffer
	if err := ipt.runWithOutput(cmd, &stdout); err != nil {
		return nil, err
	}

	UnprocessedChains := strings.Split(stdout.String(), "\n")
	var chains []string
	for _, ch := range UnprocessedChains {
		if strings.HasPrefix(ch, "-P") || strings.HasPrefix(ch, "-N") {
			chains = append(chains, strings.Fields(ch)[1])
		} else {
			break
		}
	}
	return chains, nil
}

func (ipt *IPTables) Append(table, chain string, rule ...string) error {
	cmd := append([]string{"-t", table, "-A", chain}, rule...)
	return ipt.runWithOutput(cmd, nil)
}
