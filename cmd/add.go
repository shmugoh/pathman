package cmd

import (
	"fmt"
	"strings"

	"pathman/platform"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add --folder [folder] --path [path]",
	Short: "Add folder to given PATH Key",
	Args:  cobra.MaximumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := platform.GET_ENV(pathDestination, pathKey)
		if err != nil {
			return err
		}

		fmt.Printf("Adding value to variable %s...\n", pathKey)
		pathSlice := strings.Split(path, ";")
		pathSlice = append(pathSlice, folderInput)

		for i, p := range pathSlice {
			pathValue.WriteString(p)
			if i != len(pathSlice)-1 {
				pathValue.WriteString(";")
			}
		}
		return nil
	},
}
