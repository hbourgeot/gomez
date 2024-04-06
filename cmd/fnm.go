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
	Long: `Install Fast Node Manager from its website and configure it properly.

Gomez will install it fnm from https://fnm.vercel.app/install

`,
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

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fnmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
