/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/hbourgeot/gomez/helpers"

	"github.com/spf13/cobra"
)

// nvmCmd represents the nvm command
var nvmCmd = &cobra.Command{
	Use:   "nvm",
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
		lts, _ := cmd.Flags().GetBool("lts")

		version, _ = cmd.Flags().GetString("version")

		if lts && version == "" {
			fmt.Println("Installing the latest LTS Node.js version (20.12.1)")
			version = "20.12.1"
		} else {
			fmt.Println("Installing Node.js version", version)
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
		}

		// Call the function to install fnm
		err := helpers.InstallNvm(shell, sourceFile, version)
		if err != nil {
			fmt.Println("Error installing nvm")
			fmt.Println(err)
			return
		}

		fmt.Println("Process finished.")
	},
}

func init() {
	rootCmd.AddCommand(nvmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nvmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nvmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	nvmCmd.Flags().BoolP("lts", "l", true, "Install the latest LTS Node.js version")
	nvmCmd.Flags().StringP("version", "v", "", "Install the Node.js version specified")
}
