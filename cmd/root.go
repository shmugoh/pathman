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

	pathValue       strings.Builder
	pathDestination [2]interface{}
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
				return fmt.Errorf("error finding executable: %v", err)
			}
			folderInput = filepath.Dir(executable)
		}

		// Sets PATH Location in Registry
		token := windows.GetCurrentProcessToken()
		switch isElevated := token.IsElevated(); isElevated {
		case true:
			pathDestination[0] = registry.LOCAL_MACHINE
			pathDestination[1] = `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`
		case false:
			warning := fmt.Errorf("WARNING: %v", "User-level access only. Elevated shell needed for system variables.\n")
			fmt.Fprintf(os.Stderr, "%v\n", warning)
			pathDestination[0] = registry.CURRENT_USER
			pathDestination[1] = `Environment`
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
			return err
		}
		return nil
	},

	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		err := setEnv(pathDestination, pathKey, pathValue.String())
		if err != nil {
			return err
		}
		return nil
	},
}

func getEnv(pathDest [2]interface{}, pathKey string) (string, error) {
	k, err := registry.OpenKey(pathDest[0].(registry.Key), pathDest[1].(string), registry.QUERY_VALUE|registry.WRITE)
	if err != nil {
		return "", fmt.Errorf("error opening the key: %v", err)
	}
	defer k.Close()

	// Obtains Key Value
	s, _, err := k.GetStringValue(pathKey)
	if err == nil {
		fmt.Printf("Getting value of variable %v...\n", pathKey)
		return s, nil
	}

	// If Key does not exist, set it with an empty string
	fmt.Printf("Appending variable %s to environment...\n", pathKey)
	err = k.SetExpandStringValue(pathKey, "")
	if err != nil {
		return "", fmt.Errorf("error setting the value: %v", err)
	}

	// Retrieve Key Value again
	s, _, err = k.GetStringValue(pathKey)
	if err != nil {
		return "", fmt.Errorf("error retrieving the value: %v", err)
	}

	return s, nil
}

func setEnv(pathDest [2]interface{}, pathKey string, envValue string) error {
	fmt.Printf("Writing new values to variable %s...", pathKey)
	k, err := registry.OpenKey(pathDest[0].(registry.Key), pathDest[1].(string), registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("error opening the key: %v", err)
	}
	defer k.Close()

	err = k.SetStringValue(pathKey, envValue)
	if err != nil {
		return fmt.Errorf("error setting the value: %v", err)
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
