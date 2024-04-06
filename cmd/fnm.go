/*
Copyright © 2024 Henrry Bourgeot <henrrybrgt@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fnmCmd represents the fnm command
var fnmCmd = &cobra.Command{
	Use:   "fnm",
	Short: "Install Fast Node Manager (fnm)",
	Long: `Gomez will install fnm from https://fnm.vercel.app/install, by default, it will configure the path environment variable
in the .profile file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fnm called")
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
