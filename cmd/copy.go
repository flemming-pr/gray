package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy [destination]",
	Short: "Copies the gray client to a project",
	Long:  `This command will copy the client to a project and will add the client to the .gitignore. Add the root path to your desired project.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		destination := args[0]
		fmt.Printf("Copy gray client to %s\n", destination)

		err := copyDir("gray", destination)
		if err != nil {
			fmt.Printf("Error copying gray client: %v\n", err)
		}

		err = updateGitignore(destination)
		if err != nil {
			fmt.Printf("Error updating .gitignore: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
}

func copyDir(src string, dst string) error {
	dst += "/gray"
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(path, dstPath)
	})
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return dstFile.Sync()
}

func updateGitignore(dst string) error {
	gitignorePath := filepath.Join(dst, ".gitignore")
	file, err := os.OpenFile(gitignorePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "gray" {
			return nil
		}
	}

	if _, err := file.WriteString("\ngray\n"); err != nil {
		return err
	}

	return nil
}
