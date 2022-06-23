package internal

import (
	"os"
	"strings"

	"github.com/itrepablik/itrlog"
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func SaveBlocklist(blocklist string) (bool, error) {

	// save blocklist
	b, err := os.Create(config.BlocklistFile)
	if err != nil {
		itrlog.Error(err)
		return false, err
	}
	defer b.Close()
	nw := strings.Join(strings.Split(blocklist, ","), ")|(")
	nw = "(" + nw + ")"
	if _, err = b.WriteString(pkg.Encrypt(nw, config.Key)); err != nil {
		itrlog.Error(err)
		return false, err
	}

	return true, nil
}
