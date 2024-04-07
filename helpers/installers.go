package helpers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func InstallFnm(shell, sourceFile, version string) error {
	cmd := exec.Command("curl", "-fsSL", "https://fnm.vercel.app/install")
	shellCmd := exec.Command(shell)

	shellCmd.Stdin, _ = cmd.StdoutPipe()

	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	err = shellCmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	err = shellCmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("fnm installed successfully")

	if shell != "bash" {
		homedir, _ := os.UserHomeDir()
		fnmPath := fmt.Sprintf("# fnm\nexport PATH=\"%s/.local/share/fnm:$PATH\"\neval `fnm env`\n\n", homedir)

		fmt.Println("Adding fnm to PATH")
		sourceFileDir := homedir + "/" + strings.Replace(sourceFile, "~/", "", 1)

		// Open the file in read mode
		file, err := os.Open(sourceFileDir)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer file.Close()

		// Check if the file already contains the fnm path
		scanner := bufio.NewScanner(file)
		inserted := false
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "# fnm") {
				fmt.Println("fnm already added to PATH")
				inserted = true
				break
			}
		}

		// If the string doesn't exist, open the file in append mode and add the string
		if !inserted {
			file, err := os.OpenFile(sourceFileDir, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
			defer file.Close()

			// Create a buffered writer from the file
			writer := bufio.NewWriter(file)

			// Write the text to the file
			_, err = writer.WriteString(fnmPath)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}

			// Flush the buffered writer to ensure all data is written to the file
			err = writer.Flush()
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		}
	}

	fmt.Println("Installing Node.js version", version)
	cmd = exec.Command(shell, "-c", "source "+sourceFile+" && fnm install "+version)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("Node.js version installed successfully")
	return nil
}

func InstallPyenv(shell, sourceFile, version string) error {
	cmd := exec.Command("curl", "https://pyenv.run")
	shellCmd := exec.Command(shell)

	shellCmd.Stdin, _ = cmd.StdoutPipe()

	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr

	// Start the curl command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Start the shell command
	err = shellCmd.Start()
	if err != nil {
		return err
	}

	// Wait for both commands to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	err = shellCmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("pyenv installed successfully")

	fmt.Println("Adding pyenv to PATH")

	pyenvPath := "export PYENV_ROOT=\"$HOME/.pyenv\"\n[[ -d $PYENV_ROOT/bin ]] && export PATH=\"$PYENV_ROOT/bin:$PATH\"\neval \"$(pyenv init -)\"\n\n"

	homedir, _ := os.UserHomeDir()
	sourceFileDir := homedir + "/" + strings.Replace(sourceFile, "~/", "", 1)

	// Open the file in read mode
	file, err := os.Open(sourceFileDir)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer file.Close()

	// Check if the file already contains the pyenv path
	scanner := bufio.NewScanner(file)
	inserted := false
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "PYENV_ROOT") {
			fmt.Println("pyenv already added to PATH")
			inserted = true
			break
		}
	}

	// If the string doesn't exist, open the file in append mode and add the string
	if !inserted {
		file, err := os.OpenFile(sourceFileDir, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer file.Close()

		// Create a buffered writer from the file
		writer := bufio.NewWriter(file)

		// Write the text to the file
		_, err = writer.WriteString(pyenvPath)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Flush the buffered writer to ensure all data is written to the file
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
	}

	fmt.Println("Refreshing changes in terminal with source")
	cmd = exec.Command(shell, "-c", "source "+sourceFile+" && pyenv install "+version)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("Python version installed successfully")
	return nil
}

func InstallNvm(shell, sourceFile, version string) error {
	cmd := exec.Command("curl", "-o-", "https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh")
	shellCmd := exec.Command(shell)

	shellCmd.Stdin, _ = cmd.StdoutPipe()

	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr

	// Start the curl command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Start the shell command
	err = shellCmd.Start()
	if err != nil {
		return err
	}

	// Wait for both commands to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	err = shellCmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("nvm installed successfully")
	fmt.Println("Installing Node.js version", version)
	cmd = exec.Command(shell, "-c", "source "+sourceFile+" && nvm install "+version)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("Node.js version installed successfully")
	return nil
}

func InstallSdkman(shell, sourceFile, version string) error {
	cmd := exec.Command("curl", "-s", "https://get.sdkman.io")
	shellCmd := exec.Command(shell)

	shellCmd.Stdin, _ = cmd.StdoutPipe()

	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr

	// Start the curl command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Start the shell command
	err = shellCmd.Start()
	if err != nil {
		return err
	}

	// Wait for both commands to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	err = shellCmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("SDKMan installed successfully")

	cmd = exec.Command(shell, "-c", `source "$HOME/.sdkman/bin/sdkman-init.sh" && sdk install java `+version+"-tem")
	fmt.Println("Installing Java version", version, "Temurin SDK")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("Java version installed successfully")
	return nil
}

func InstallCargo(shell, sourceFile string) error {
	cmd := exec.Command("curl", "https://static.rust-lang.org/rustup.sh", "-sSf")
	shellCmd := exec.Command(shell)

	shellCmd.Stdin, _ = cmd.StdoutPipe()

	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr

	// Start the curl command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Start the shell command
	err = shellCmd.Start()
	if err != nil {
		return err
	}

	// Wait for both commands to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	err = shellCmd.Wait()
	if err != nil {
		return err
	}

	fmt.Println("Cargo installed successfully")
	fmt.Println("Refreshing changes in terminal with source")
	cmd = exec.Command(shell, "-c", "source "+sourceFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Changes refreshed successfully")
	return nil
}
