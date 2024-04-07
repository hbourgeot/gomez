/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/hbourgeot/gomez/helpers"

	"github.com/spf13/cobra"
)

// cargoCmd represents the cargo command
var cargoCmd = &cobra.Command{
	Use:   "cargo",
	Short: "Cargo, the package manager for Rust",
	Long: `Cargo is the package manager for Rust. It is used to compile, test, and run Rust programs.

Gomez will install Cargo from https://rustup.rs/ and configure the path environment variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		// variables
		var shell, sourceFile string

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
		err := helpers.InstallCargo(shell, sourceFile)
		if err != nil {
			fmt.Println("Error installing cargo")
			fmt.Println(err)
			return
		}

		fmt.Println("Process finished.")
	},
}

func init() {
	rootCmd.AddCommand(cargoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cargoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cargoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
