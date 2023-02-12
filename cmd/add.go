package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add --folder [folder] --path [path]",
	Short: "Add folder to given PATH Key",
	Args:  cobra.MaximumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := getEnv(pathDest, pathKey)
		if err != nil {
			fmt.Errorf(err.Error())
		}

		fmt.Printf("Editing value of %s...\n", pathKey)
		pathSlice := strings.Split(path, ";")
		pathSlice = append(pathSlice, folderInput)
		var buff strings.Builder
		for i, p := range pathSlice {
			buff.WriteString(p)
			if i != len(pathSlice)-1 {
				buff.WriteString(";")
			}
		}

		setEnv(pathDest, pathKey, buff.String())
		return nil
	},
}
