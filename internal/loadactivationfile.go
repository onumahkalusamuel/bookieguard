package internal

import (
	"encoding/json"
	"os"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func LoadActivationFile() (bool, config.BodyStructure) {

	f, err := os.ReadFile(config.ActivationFile)
	if err != nil || len(string(f)) < 36 {
		return false, config.BodyStructure{}
	}

	decrypted, err := pkg.Decrypt(string(f), config.Key)
	if err != nil {
		return false, config.BodyStructure{}
	}

	var unmarshalled config.BodyStructure

	json.Unmarshal([]byte(decrypted), &unmarshalled)

	return true, unmarshalled

}
