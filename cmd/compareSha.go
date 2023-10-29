/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// compareShaCmd represents the compareSha command
var compareShaCmd = &cobra.Command{
	Use:   "compareSha",
	Short: "verify your docker images",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var yourDockerImage string = args[0]
		var configFilePath string = args[1]
		fmt.Println("compareSha called")

	},
}

func init() {
	rootCmd.AddCommand(compareShaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compareShaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareShaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
