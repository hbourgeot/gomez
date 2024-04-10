/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/hbourgeot/gomez/colors"
	"github.com/hbourgeot/gomez/helpers"
	"github.com/spf13/cobra"
	"strings"
)

// sdkmCmd represents the sdkm command
var sdkmCmd = &cobra.Command{
	Use:   "sdkm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// variables
		var version, shell, sourceFile string
		// version
		version, _ = cmd.Flags().GetString("version")
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
			validLts := strings.HasSuffix(version, "17") ||
				strings.HasSuffix(version, "11") ||
				strings.HasSuffix(version, "21") ||
				strings.HasSuffix(version, "8")

			if !validLts {
				fmt.Println("The version provided", colors.Yellow, "is not a LTS version.", colors.Reset, "Please provide a valid LTS version.")
				return
			}

			fmt.Println("Installing the LTS Java version", version)
		} else {
			fmt.Println("Installing Java version", version)
		}

		forZsh, _ := cmd.Flags().GetBool("zsh")
		forBash, _ := cmd.Flags().GetBool("bash")
		forFish, _ := cmd.Flags().GetBool("fish")

		if forZsh {
			shell = "zsh"
			sourceFile = "$HOME/.zshrc"
		} else if forFish {
			shell = "fish"
			sourceFile = "$HOME/.config/fish/config.fish"
		} else if forBash {
			shell = "bash"
			sourceFile = "$HOME/.bashrc"
		} else {
			shell = "bash"
			sourceFile = "~/.bashrc"

			fmt.Println("Shell not specified, using " + colors.Green + "bash" + colors.Reset + " as default.")
		}

		// Call the function to install fnm
		err := helpers.InstallSdkman(shell, sourceFile, version)
		if err != nil {
			fmt.Println("Error installing sdkman")
			fmt.Println(err)
			return
		}

		fmt.Println("Process finished.")
	},
}

func init() {
	rootCmd.AddCommand(sdkmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sdkmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sdkmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	sdkmCmd.Flags().BoolP("lts", "l", true, "Install the latest LTS Node.js version")
	sdkmCmd.Flags().StringP("version", "v", "", "Install the Node.js version specified")
}
