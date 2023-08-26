//go:build windows

package platform

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func SET_PATH(pathDestination [2]interface{}, pathKey string) (string, error) {
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
	return pathKey, nil
}

func GET_ENV(pathDest [2]interface{}, pathKey string) (string, error) {
	// Open OS PATH Key
	k, err := registry.OpenKey(pathDest[0].(registry.Key), pathDest[1].(string), registry.QUERY_VALUE|registry.WRITE)
	if err != nil {
		return "", fmt.Errorf("error opening the key: %v", err)
	}
	defer k.Close()

	// Obtains Picked PATH Name from Key
	s, _, err := k.GetStringValue(pathKey)
	if err == nil {
		fmt.Printf("Getting value of variable %v...\n", pathKey)
		return s, nil
	}

	// Creates New PATH if non-existant
	fmt.Printf("Appending variable %s to environment...\n", pathKey)
	err = k.SetExpandStringValue(pathKey, "")
	if err != nil {
		return "", fmt.Errorf("error setting the value: %v", err)
	}

	// Obtains Picked PATH Name from Key
	s, _, err = k.GetStringValue(pathKey)
	if err != nil {
		return "", fmt.Errorf("error retrieving the value: %v", err)
	}

	// Returns Picked PATH Value
	return s, nil
}

func SET_ENV(pathDest [2]interface{}, pathKey string, envValue string) error {
	// Write Changes to PATH
	fmt.Printf("Writing new values to variable %s...", pathKey)

	// Open PATH Key
	k, err := registry.OpenKey(pathDest[0].(registry.Key), pathDest[1].(string), registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("error opening the key: %v", err)
	}
	defer k.Close()

	// Write changes to Picked PATH from PATH key
	err = k.SetStringValue(pathKey, envValue)
	if err != nil {
		return fmt.Errorf("error setting the value: %v", err)
	}
	return nil
}