/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/hbourgeot/gomez/colors"
	"github.com/hbourgeot/gomez/helpers"

	"github.com/spf13/cobra"
)

// pyenvCmd represents the pyenv command
var pyenvCmd = &cobra.Command{
	Use:   "pyenv",
	Short: "Install Python Version Management",
	Long: `Install pyenv from its website https://pyenv.run.

Provide a version for install, default is 3.9. By default, the path environment variable is configured on .profile file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// variables
		var version, shell, sourceFile string
		// version
		version, _ = cmd.Flags().GetString("version")
		if version == "" {
			version = "3.10"
		}

		// check if the version is valid (only numbers and dots)
		if !helpers.IsValidVersion(version) {
			fmt.Println(colors.Red + "Invalid version!" + colors.Reset + " Please provide a valid version number. For example: " + colors.Green + "3.10 or 3.10.4" + colors.Reset)
			return
		}

		if !helpers.VersionExists(version, "pyenv") {
			fmt.Println(colors.Red + "Version does not exist (yet)!" + colors.Reset + " Please provide a valid version number. For example: " + colors.Green + "3.10 or 3.10.4" + colors.Reset)
			return
		}

		fmt.Println("Installing Python version", version)

		forZsh, _ := cmd.Flags().GetBool("zsh")
		forBash, _ := cmd.Flags().GetBool("bash")
		forFish, _ := cmd.Flags().GetBool("fish")

		if forZsh {
			shell = "zsh"
			sourceFile = "~/.zshrc"
		} else if forFish {
			shell = "fish"
			sourceFile = "~/.config/fish/config.fish"
		} else if forBash {
			shell = "bash"
			sourceFile = "~/.bashrc"
		} else {
			shell = "bash"
			sourceFile = "~/.bashrc"

			fmt.Println("Shell not specified, using " + colors.Green + "bash" + colors.Reset + " as default.")
		}

		// Call the function to install fnm
		err := helpers.InstallPyenv(shell, sourceFile, version)
		if err != nil {
			fmt.Println("Error installing pyenv")
			fmt.Println(err)
			return
		}

		fmt.Println("Process finished.")
	},
}

func init() {
	rootCmd.AddCommand(pyenvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pyenvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	pyenvCmd.Flags().StringP("version", "v", "", "Install the Node.js version specified")
}
