/*
Copyright © 2024 Henrry Bourgeot <henrrybrgt@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gmz",
	Short: "Gomez, your installer helper",
	Long: `
Gomez, your loyal helper for install some programming tools, like: Fast Node Manager (fnm), Pyenv,
Node Version Manager (nvm), Cargo, SDKMan and Make, for now.

Why I'm developing this CLI? For automatize the installations of the most programming tools needed
for programmers. With a few keys press. Don't go to the github page of any of these, Gomez can help
you with that. Only write the tool that you need and Gomez will:

1. Install the tool.
2. Configure the tool for the shell of your choice.
3. Install the latest stable version of the programming language (or provide a version).

"Gomez, I'm to serve him, my lord."`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gomez.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
