package cmd

import (
	"fmt"
	"pathman/platform"
	"strings"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove --folder [folder] --path [path]",
	Short: "Remove folder to given PATH Key",
	Args:  cobra.MaximumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := platform.GET_ENV(pathDestination, pathKey)
		if err != nil {
			return err
		}

		fmt.Printf("Removing value from variable %s...\n", pathKey)
		pathSlice := strings.Split(path, ";")
		pathSlice = append(pathSlice, folderInput)

		for i, p := range pathSlice {
			if p != folderInput {
				pathValue.WriteString(p)
				if i != len(pathSlice)-1 {
					pathValue.WriteString(";")
				}
			}
		}
		return nil
	},
}
