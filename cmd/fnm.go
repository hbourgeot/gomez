/*
Copyright © 2024 Henrry Bourgeot <henrrybrgt@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/hbourgeot/gomez/colors"
	"github.com/hbourgeot/gomez/helpers"
	"strconv"

	"github.com/spf13/cobra"
)

// fnmCmd represents the fnm command
var fnmCmd = &cobra.Command{
	Use:   "fnm",
	Short: "Install Fast Node Manager (fnm)",
	Long: `Gomez will install fnm from https://fnm.vercel.app/install, by default, it will configure the path environment variable
in the .profile file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// variables
		var version, shell, sourceFile string
		// version
		lts, _ := cmd.Flags().GetBool("lts")

		if !helpers.IsValidVersion(version) {
			fmt.Println(colors.Red, "Invalid version!", colors.Reset, "Please provide a valid version number. For example:", colors.Yellow, "16.10 or 20.0.4", colors.Reset)
			return
		}

		if !helpers.VersionExists(version, "nvm") {
			fmt.Println(colors.Red, "Version does not exist (yet)!", colors.Reset, "Please provide a valid version number. For example:", colors.Yellow, "16.10 or 20.0.4", colors.Reset)
			return
		}

		if lts && version == "" {
			fmt.Println("Installing the latest LTS Node.js version (20.12.1)")
			version = "20.12.1"
		} else if lts && version != "" {
			localVersion, _ := strconv.Atoi(version[:2])
			if localVersion%2 != 0 {
				fmt.Println("The version provided", colors.Yellow, "is not a LTS version.", colors.Reset, "Please provide a valid LTS version.")
				return
			}

			fmt.Println("Installing the LTS Node.js version", version)
		} else {
			fmt.Println("Installing Node.js version", version)
		}

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
		err := helpers.InstallFnm(shell, sourceFile, version)
		if err != nil {
			fmt.Println("Error installing fnm")
			fmt.Println(err)
			return
		}

		fmt.Println("Process finished.")
	},
}

func init() {
	rootCmd.AddCommand(fnmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fnmCmd.PersistentFlags().String("foo", "", "A help for foo")

	fnmCmd.Flags().BoolP("lts", "l", true, "Install the latest LTS Node.js version")
	fnmCmd.Flags().StringP("version", "v", "", "Install the Node.js version specified")
}
