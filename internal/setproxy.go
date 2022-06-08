package internal

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
)

func SetProxy() {
	ct.Foreground(ct.Cyan, false)
	fmt.Println("Setting proxy...")
	ct.ResetColor()

}
