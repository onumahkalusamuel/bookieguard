package internal

import (
	"context"
	"fmt"
	"net"
	"regexp"

	"github.com/elazarl/goproxy"
	"github.com/onumahkalusamuel/bookieguard/config"
)

func StartServer() {

	if config.PROXY_SERVER_HANDLE.Addr != "" {
		_, err := StopServer()
		if err != nil {
			fmt.Println("Unable to shutdown: ", err)
		}
	}

	// set up the proxy
	// go SetProxy()

	// continue
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

	config.PROXY_SERVER_HANDLE.Addr = net.JoinHostPort(config.PROXY_HOST, config.PROXY_PORT)
	config.PROXY_SERVER_HANDLE.Handler = proxy

	// launch
	err := config.PROXY_SERVER_HANDLE.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

	config.PROXY_SERVER_HANDLE.Shutdown(context.Background())

}
