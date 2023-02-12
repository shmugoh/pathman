package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

var (
	folderInput string
	pathKey     string

	path     string
	pathDest [2]interface{}
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
			executable, err := os.Executable()
			if err != nil {
				return fmt.Errorf(err.Error())
			}
			folderInput = filepath.Dir(executable)
		}

		// Sets PATH Location in Registry
		token := windows.GetCurrentProcessToken()
		switch isElevated := token.IsElevated(); isElevated {
		case true:
			pathDest[0] = registry.LOCAL_MACHINE
			pathDest[1] = `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`
		case false:
			pathDest[0] = registry.CURRENT_USER
			pathDest[1] = `Environment`
		}
		token.Close()

		// Sets PATH Key
		pathKey = strings.ToLower(pathKey)

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		// Runs `addCmd` if no command is given
		err := addCmd.RunE(cmd, args)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		return nil
	},
}

func getEnv(pathDest [2]interface{}, pathKey string) (string, error) {
	k, err := registry.OpenKey(pathDest[0].(registry.Key), pathDest[1].(string), registry.QUERY_VALUE)
	if err != nil {
		fmt.Errorf("error opening the key: %s", err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue(pathKey)
	if err == nil {
		fmt.Printf("Getting value of %s...\n", pathKey)
		return s, nil
	}

	// If the key does not exist, set it with an empty string
	fmt.Printf("Appending %s to environment...\n", pathKey)
	err = k.SetExpandStringValue(pathKey, "")
	if err != nil {
		fmt.Errorf("error setting the value: %s", err)
	}

	// Retrieve the value again
	s, _, err = k.GetStringValue(pathKey)
	if err != nil {
		fmt.Errorf("error retrieving the value: %s", err)
	}

	return s, nil
}

func setEnv(pathDest [2]interface{}, pathKey string, envValue string) error {
	fmt.Printf("Writing new values to %s...", pathKey)
	k, err := registry.OpenKey(pathDest[0].(registry.Key), pathDest[1].(string), registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.SetStringValue(pathKey, envValue)
	if err != nil {
		return err
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
