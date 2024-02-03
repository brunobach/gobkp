package create

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/brunobach/gobkp/internal/pkg/helper"
	"github.com/spf13/cobra"
)

var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup specified files and directories",
	Run: func(cmd *cobra.Command, args []string) {
		configFile := "backup.cfg"
		zipFilename := "backup.zip"
		logFilename := "backup.log"

		files, excludes, err := helper.ReadConfig(configFile)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening log file: %v", err)
		}
		defer logFile.Close()

		log.SetOutput(logFile)

		if err := createZip(zipFilename, files, excludes); err != nil {
			fmt.Println("Error creating zip file:", err)
			return
		}

		fmt.Println("Backup completed successfully!")
	},
}

func createZip(zipFilename string, files []string, excludes []string) error {
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	for _, file := range files {
		matches, err := filepath.Glob(file)
		if err != nil {
			log.Println("Error matching files:", err)
			continue
		}
		for _, match := range matches {
			if !isExcluded(match, excludes) {
				if err := addToZip(archive, match); err != nil {
					log.Printf("Error adding file %s to zip: %v\n", match, err)
				}
			}
		}
	}

	return nil
}

func addToZip(archive *zip.Writer, file string) error {
	info, err := os.Stat(file)
	if err != nil {
		return err
	}

	homeDir := os.Getenv("HOME")

	if info.IsDir() {
		err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			relPath, err := filepath.Rel(homeDir, path)
			if err != nil {
				return err
			}
			header.Name = relPath
			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}
			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}
			if !info.IsDir() {
				src, err := os.Open(path)
				if err != nil {
					return err
				}
				defer src.Close()
				_, err = io.Copy(writer, src)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		if strings.HasPrefix(file, "~/.") {
			file = strings.Replace(file, "~", homeDir, 1)
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = filepath.Base(file)
		header.Method = zip.Deflate
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		src, err := os.Open(file)
		if err != nil {
			return err
		}
		defer src.Close()
		_, err = io.Copy(writer, src)
		if err != nil {
			return err
		}
	}

	return nil
}

func isExcluded(file string, excludes []string) bool {
	for _, pattern := range excludes {
		if matched, _ := filepath.Match(pattern, file); matched {
			return true
		}
	}
	return false
}
