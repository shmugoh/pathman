//go:build linux

package platform

func SET_PATH(pathDestination [2]interface{}, pathKey string) (string, error) {
	// Sets PATH Location in Registry

	// Sets PATH Key
	return "s", nil
}

// func getEnv(pathDest [2]interface{}, pathKey string) (string, error) {
func GET_ENV(pathDest [2]interface{}, pathKey string) ([2]interface{}, string, error) {
	// Open OS PATH Key

	// Obtains Picked PATH Name from Key

	// Creates New PATH if non-existant

	// Obtains Picked PATH Name from Key
	return {{}, {}}, "s", nil
}

func SET_ENV(pathDest [2]interface{}, pathKey string, envValue string) error {
	// Write Changes to PATH
	
	// Open PATH Key

	// Write changes to Picked PATH from PATH key

	return nil
}