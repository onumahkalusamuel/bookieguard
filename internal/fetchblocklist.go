package internal

import (
	"os"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func FetchBlockList() string {
	f, err := os.ReadFile(config.BlocklistFile)
	if err != nil {
		return "*"
	}
	return pkg.Decrypt(string(f), config.Key)
}
