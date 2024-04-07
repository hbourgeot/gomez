package helpers

import (
	"fmt"
	"os"
	"testing"
)

func TestInstallFnm(t *testing.T) {
	type args struct {
		shell      string
		sourceFile string
		version    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: string("Test fnm installation with a version with a v prefix"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("v1.22.0"),
			},
			wantErr: true,
		},
		{
			name: string("Test fnm installation with a invalid version"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("1.22.0"),
			},
			wantErr: true,
		},
		{
			name: string("Test fnm installation with a valid version"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("18"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InstallFnm(tt.args.shell, tt.args.sourceFile, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("InstallFnm() error = %v, wantErr %v", err, tt.wantErr)
			}

			deleteAfterTest("fnm")
		})
	}
}

func TestInstallPyenv(t *testing.T) {
	type args struct {
		shell      string
		sourceFile string
		version    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: string("Test pyenv installation with a version with a v prefix"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("v1.22.0"),
			},
			wantErr: true,
		},
		{
			name: string("Test pyenv installation with a invalid version"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("1.22.0"),
			},
			wantErr: true,
		},
		{
			name: string("Test pyenv installation with a valid version"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("3.8"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InstallPyenv(tt.args.shell, tt.args.sourceFile, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("InstallPyenv() error = %v, wantErr %v", err, tt.wantErr)
			}

			deleteAfterTest("pyenv")
		})
	}
}

func TestInstallNvm(t *testing.T) {
	type args struct {
		shell      string
		sourceFile string
		version    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: string("Test nvm installation with a version with a v prefix"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("v1.22.0"),
			},
			wantErr: true,
		},
		{
			name: string("Test nvm installation with a invalid version"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("1.22.0"),
			},
			wantErr: true,
		},
		{
			name: string("Test nvm installation with a valid version"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("18"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InstallNvm(tt.args.shell, tt.args.sourceFile, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("InstallNvm() error = %v, wantErr %v", err, tt.wantErr)
			}

			deleteAfterTest("nvm")
		})
	}
}

func TestInstallSdkman(t *testing.T) {
	type args struct {
		shell      string
		sourceFile string
		version    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: string("Test sdk installation with a version with a v prefix"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("v1.22.0"),
			},
			wantErr: true,
		},
		{
			name: string("Test sdk installation with a invalid version"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("1.22.0"),
			},
			wantErr: true,
		},
		{
			name: string("Test sdk installation with a valid version"),
			args: args{
				shell:      string("bash"),
				sourceFile: string("~/.bashrc"),
				version:    string("21"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InstallSdkman(tt.args.shell, tt.args.sourceFile, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("InstallSdkman() error = %v, wantErr %v", err, tt.wantErr)
			}

			deleteAfterTest("sdkman")
		})
	}
}

func deleteAfterTest(tool string) {
	var path string
	homePath, _ := os.UserHomeDir()

	switch tool {
	case "fnm":
		path = fmt.Sprintf("%s/.local/share/fnm/fnm", homePath)
	case "pyenv":
		path = fmt.Sprintf("%s/.pyenv", homePath)
	case "nvm":
		path = fmt.Sprintf("%s/.nvm", homePath)
	case "sdkman":
		path = fmt.Sprintf("%s/.sdkman", homePath)
	}

	os.RemoveAll(path)
}
