package internal

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/onumahkalusamuel/bookieguard/config"
)

func StartServer() {
	// test if port is available
	portAvailable := false

	conn, err := net.DialTimeout(
		"tcp",
		net.JoinHostPort(config.PROXY_HOST, config.PROXY_PORT),
		2*time.Second,
	)
	if err == nil {
		portAvailable = true
	}
	conn.Close()

	time.Sleep(2 * time.Second)

	if portAvailable {

		var hostRegex *regexp.Regexp = CompileRegexp()

		proxy := goproxy.NewProxyHttpServer()

		proxy.OnRequest(goproxy.UrlMatches(regexp.MustCompile("^.*$"))).
			HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
				if hostRegex.MatchString(host) == true {
					return goproxy.RejectConnect, host
				}

				go SaveHost(host)
				return goproxy.OkConnect, host
			})

		config.PROXY_SERVER_HANDLE.ListenAndServe()

		config.PROXY_SERVER_HANDLE = &http.Server{
			Addr:    net.JoinHostPort(config.PROXY_HOST, config.PROXY_PORT),
			Handler: proxy,
		}
		// launch
		err := config.PROXY_SERVER_HANDLE.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}

}
