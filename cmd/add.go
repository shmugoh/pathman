package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	// Use:   "execute [URL] [message]\n  dishook execute [URL]",
	// Short: "Sends message and/or arguments to Discord",
	// Args:  cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		err := addFunc()
		if err != nil {
			return err
		}
		return nil
	},
}

func addFunc() error {
	path = getEnv(pathDest, pathKey)

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
}
