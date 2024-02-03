package restore

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/brunobach/gobkp/internal/pkg/helper"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore files and directories from a zip file",
	Run: func(cmd *cobra.Command, args []string) {
		zipFilename := "backup.zip"
		configFile := "backup.cfg"

		files, _, err := helper.ReadConfig(configFile)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		if err := extractFromZip(zipFilename, files); err != nil {
			fmt.Println("Error restoring files:", err)
			return
		}

		fmt.Println("Restore completed successfully!")
	},
}

func extractFromZip(zipFilename string, files []string) error {
	zipFile, err := zip.OpenReader(zipFilename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	extractMap := make(map[string]bool)
	for _, file := range files {
		matches, err := filepath.Glob(file)
		if err != nil {
			return err
		}
		for _, match := range matches {
			extractMap[match] = true
		}
	}

	for _, file := range zipFile.File {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		if _, ok := extractMap[file.Name]; ok {
			destPath := filepath.Join(".", file.Name)
			if file.FileInfo().IsDir() {
				os.MkdirAll(destPath, os.ModePerm)
			} else {
				dest, err := os.Create(destPath)
				if err != nil {
					return err
				}
				defer dest.Close()

				_, err = io.Copy(dest, src)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
