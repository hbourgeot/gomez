package helpers

import (
	"fmt"
	"os"
	"os/exec"
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
	fmt.Println("Refreshing changes in terminal with source")
	cmd = exec.Command(shell, "-c", "source "+sourceFile)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Changes refreshed successfully")
	fmt.Println("Installing Node.js version", version)
	cmd = exec.Command("fnm", "install", version)

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
	fmt.Println("Refreshing changes in terminal with source")
	cmd = exec.Command(shell, "-c", "source "+sourceFile)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Changes refreshed successfully")
	fmt.Println("Installing Python version", version)
	cmd = exec.Command("pyenv", "install", version)

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
	fmt.Println("Refreshing changes in terminal with source")
	cmd = exec.Command("source", sourceFile)

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

	fmt.Println("Changes refreshed successfully")
	fmt.Println("Installing Node.js version", version)
	cmd = exec.Command(shell, "-c", "source "+sourceFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
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
