package internal

import (
	"fmt"
	"regexp"
)

func CompileRegexp() *regexp.Regexp {
	blocklist := FetchBlockList()
	fmt.Println("compiling")
	// nw := strings.Join(blocklist, ")|(")
	return regexp.MustCompile(fmt.Sprintf(".*(%s).*$", blocklist))
}
