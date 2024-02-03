package cmd

import (
	"fmt"
	"os"

	"github.com/brunobach/gobkp/internal/command/create"
	"github.com/brunobach/gobkp/internal/command/restore"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobkp",
	Short: "A simple CLI for backup and restore",
	Long: `gobkp is a CLI application to perform backup and restore operations.
	It can backup specified files and directories into a zip file, as well as restore files and directories from a zip file.`,
}

func init() {
	rootCmd.AddCommand(create.BackupCmd)
	rootCmd.AddCommand(restore.RestoreCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
