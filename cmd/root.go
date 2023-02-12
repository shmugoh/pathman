package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	rootCmd.AddCommand(addCmd)
}

var rootCmd = &cobra.Command{
	// Use:  "dishook [url] [message]\n  dishook [url] [flags]",
	// Args: cobra.MinimumNArgs(2),

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// if len(args) == 0 {
		// 	cmd.Help()
		// 	// return fmt.Errorf("no arguments given")
		// }
		if len(folderInput) == 0 {
			executable, err := os.Executable()
			if err != nil {
				return fmt.Errorf(err.Error())
			}
			folderInput = filepath.Dir(executable)
		}

		// Sets PATH location in Registry
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
		if pathKey == "PATH" {
			pathKey = "Path" // equivalent of registry
		}

		// fmt.Println(folderInput, pathKey)
		return nil
	},

	// if no command is parsed
	RunE: func(cmd *cobra.Command, args []string) error {
		err := addFunc()
		if err != nil {
			fmt.Errorf(err.Error())
		}
		return nil
	},
}

func getEnv(pathDest [2]interface{}, pathKey string) string {
	fmt.Printf("Getting value of %s...\n", pathKey)
	k, err := registry.OpenKey(pathDest[0].(registry.Key), pathDest[1].(string), registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	s, _, err := k.GetStringValue(pathKey)
	if err != nil {
		log.Fatal((err))
	}
	defer k.Close()
	return s
}

func setEnv(pathDest [2]interface{}, pathKey string, envValue string) error {
	fmt.Printf("Writing new values to %s...", pathKey)
	k, err := registry.OpenKey(pathDest[0].(registry.Key), pathDest[1].(string), registry.SET_VALUE)
	if err != nil {
		return err
	}
	err = k.SetStringValue(pathKey, envValue)
	if err != nil {
		return err
	}
	defer k.Close()
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
