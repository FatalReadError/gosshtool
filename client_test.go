package gosshtool

import (
	"testing"
)

func Test_Cmd(t *testing.T) {
	sshconfig := &SSHClientConfig{
		User:     "user",
		Password: "pwd",
		Host:     "11.11.22.22",
	}
	sshclient := NewSSHClient(sshconfig)
	t.Log(sshclient.Host)
	stdout, stderr, err := sshclient.Cmd("pwd")
	if err != nil {
		t.Error(err)
	}
	t.Log(stdout)
	t.Log(stderr)
}
