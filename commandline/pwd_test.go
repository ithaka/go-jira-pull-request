package commandline

import (
	"testing"
	"github.com/pkg/errors"
)

type TestRunnerPath1 struct{}
func (r TestRunnerPath1) Run(command string, args ...string) ([]byte, error) {
	failIfWrongArgs(command)
	return []byte("/this/is/my/path"), nil
}
func failIfWrongArgs(command string) {
	if command != "pwd" {
		panic("FAIL: Expected command pwd, got " + command)
	}
}

type TestRunnerPath2 struct{}
func (r TestRunnerPath2) Run(command string, args ...string) ([]byte, error) {
	failIfWrongArgs(command)
	return []byte("/this/is/my/other/path\n"), nil
}

var testError = errors.New("test")

type TestRunnerError struct{}
func (r TestRunnerError) Run(command string, args ...string) ([]byte, error) {
	failIfWrongArgs(command)
	return nil, testError
}

func TestPwd(t *testing.T) {
	tests := []struct {
		name string
		runner Runner
		output string
		err error
	}{
		{
			name:   "Should get Path",
			runner: TestRunnerPath1{},
			output:   "/this/is/my/path",
		},
		{
			name:   "Should get another Path",
			runner: TestRunnerPath2{},
			output:   "/this/is/my/other/path",
		},
		{
			name:   "Should produce an error",
			runner: TestRunnerError{},
			err: testError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner = tt.runner
			got, err := Pwd()
			if got != tt.output {
				t.Errorf("Pwd(): expected %v, got %v", tt.output, got)
			}
			if err != tt.err {
				t.Errorf("Pwd(): expected ERROR %v, got %v", tt.err, err)
			}
		})
	}
}
