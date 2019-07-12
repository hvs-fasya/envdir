package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/kami-zh/go-capturer"
)

const (
	aEnv = `123`
	bEnv = `another_val`
)

var envDir string

func TestMain(m *testing.M) {
	envDir, _ = ioutil.TempDir("../", "env_vars")
	_ = ioutil.WriteFile(envDir+"/A_ENV", []byte(aEnv), 0644)
	_ = ioutil.WriteFile(envDir+"/B_ENV", []byte(bEnv), 0644)
	os.Exit(m.Run())
}

//todo: unsufficient args
func TestRootExecute(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "simple", args: []string{envDir, `../some_prog`}, want: aEnv + " " + bEnv},
		{name: "with args", args: []string{envDir, `go`, `version`}, want: "go version " + runtime.Version() + " linux/amd64\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RootCmd.Run(nil, tt.args)
			output := capturer.CaptureStdout(func() {
				RootCmd.Run(nil, tt.args)
			})
			if tt.want != output {
				t.Errorf("RootCmd.Run output = %v, want %v", output, tt.want)
			}
		})
	}
	os.RemoveAll(envDir)
}

func captureStdout(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return buf.String()
}
