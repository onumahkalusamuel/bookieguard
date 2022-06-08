package internal

import (
	"github.com/onumahkalusamuel/bookieguard/config"
)

func CheckActivation() (bool, int) {

	// read from activation file
	loaded, unmarshalled := LoadActivationFile()
	if !loaded {
		return false, config.CHECK_ACTIVATION_READFAILURE
	}

	if unmarshalled["hashedID"] != config.HashedID || unmarshalled["activated"] != "true" {
		return false, config.CHECK_ACTIVATION_NOTACTIVATED
	}
	return true, 0
}
