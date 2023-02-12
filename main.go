package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func getEnv(pathDest [2]interface{}, pathKey string) string {
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

func main() {
	/*
		SYSTEM = SYSTEM\CurrentControlSet\Control\Session Manager\Environment
		USER = HKEY_CURRENT_USER\Environment

		"PATH" registry equivalent is "Path"
	*/

	var pathDest [2]interface{}
	token := windows.GetCurrentProcessToken()
	switch isElevated := token.IsElevated(); isElevated {
	case true:
		fmt.Println(true)
		pathDest[0] = registry.LOCAL_MACHINE
		pathDest[1] = `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`
	case false:
		fmt.Println(false)
		pathDest[0] = registry.CURRENT_USER
		pathDest[1] = `Environment`
	}
	token.Close()

	fmt.Println("Getting Path...")
	pathKey := "PATH"
	if pathKey == "PATH" {
		pathKey = "Path" // registry equivalent
	}

	fmt.Println("Getting Path Value...")
	path := getEnv(pathDest, pathKey)

	fmt.Println("Editing Path Value...")
	pathSlice := strings.Split(path, ";")
	pathSlice = append(pathSlice, "C:\\test")
	var buff strings.Builder
	for i, p := range pathSlice {
		buff.WriteString(p)
		if i != len(pathSlice)-1 {
			buff.WriteString(";")
		}
	}

	// fmt.Println(buff.String())
	fmt.Println("Writing...")
	setEnv(pathDest, pathKey, buff.String())
	fmt.Println("Done!")
}
