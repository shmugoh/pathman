package cmd

import (
	"fmt"
	"os"
	"strings"

	"pathman/platform"

	"github.com/spf13/cobra"
)

var (
	folderInput string
	pathKey     string

	pathValue       strings.Builder
	pathDestination [2]interface{}
	err             error
)

func init() {

	// add flag for folder (default is current folder)
	rootCmd.PersistentFlags().StringVarP(
		&folderInput,
		"folder",
		"f",
		"",
		"sets folder - default is curernt folder",
	)

	// add flag for path key
	rootCmd.PersistentFlags().StringVarP(
		&pathKey,
		"path",
		"p",
		"PATH",
		"sets path key to edit",
	)

	rootCmd.AddCommand(addCmd, removeCmd)
}

var rootCmd = &cobra.Command{
	Use:  "pathman",
	Args: cobra.MaximumNArgs(2),

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Obtains Current Folder if none is parsed
		if len(folderInput) == 0 {
			folderInput, err = os.Getwd()
			if err != nil {
				return fmt.Errorf("error finding executable: %v", err)
			}
		}

		// Sets PATH Location & Key
		pathKey, err = platform.SET_PATH(pathDestination, pathKey)
		if err != nil {
			return err
		}

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		// Runs `addCmd` if no command is given
		err := addCmd.RunE(cmd, args)
		if err != nil {
			return err
		}
		return nil
	},

	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		err := platform.SET_ENV(pathDestination, pathKey, pathValue.String())
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
