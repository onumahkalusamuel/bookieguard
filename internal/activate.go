package internal

import (
	"encoding/json"
	"os"

	"github.com/itrepablik/itrlog"
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func Activate(data config.BodyStructure) (bool, error) {

	store := config.BodyStructure{
		"computerName":   config.ComputerName,
		"hashedID":       config.HashedID,
		"email":          config.Email,
		"activated":      "true",
		"unlockCode":     data["unlockCode"],
		"expirationDate": data["expirationDate"],
	}

	// store configuration details
	f, err := os.Create(config.ActivationFile)
	if err != nil {
		itrlog.Error(err)
		return false, err
	}
	defer f.Close()
	marshalled, err := json.Marshal(store)
	if err != nil {
		itrlog.Error(err)
		return false, err
	}
	if _, err = f.WriteString(pkg.Encrypt(string(marshalled), config.Key)); err != nil {
		itrlog.Error(err)
		return false, err
	}

	// save blocklist
	_, err = SaveBlocklist(data["blocklist"])
	if err != nil {
		return false, err
	}

	go StartServer()

	return true, nil
}
