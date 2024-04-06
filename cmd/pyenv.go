/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pyenvCmd represents the pyenv command
var pyenvCmd = &cobra.Command{
	Use:   "pyenv",
	Short: "Install Python Version Management",
	Long: `Install pyenv from its website https://pyenv.run.

Provide a version for install, default is 3.9. By default, the path environment variable is configured on .profile file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pyenv called")
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
