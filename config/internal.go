package config

import (
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/denisbrodbeck/machineid"
)

// server things
const (
	WEB_HOST   = "localhost"
	WEB_PORT   = "8887"
	PROXY_HOST = "localhost"
	PROXY_PORT = "8888"
)

var (
	Email          string
	ActivationCode string
	UnlockCode     string
	WebBase        = "https://bookieguard.herokuapp.com/"
	ApiBase        = "https://bookieguard.herokuapp.com/api/"
	Endpoints      = map[string]string{
		"activation":       ApiBase + "activation",
		"update":           ApiBase + "update",
		"download-updates": ApiBase + "download-updates",
		"upload-hosts":     ApiBase + "upload-hosts",
		"system-status":    ApiBase + "system-status",
	}
)

// settle directory issues
var (
	s, _          = os.Executable()
	BASE_DIR      = filepath.Dir(s)
	HTTP_BASE_DIR = http.Dir(BASE_DIR + "/web/")
)

// keep a handler for opening and closing proxy connection
var PROXY_SERVER_HANDLE = new(http.Server)

// internal variable
var (
	AppKey             = "b00ki-GURd"
	HashedID, err      = machineid.ProtectedID(AppKey)
	ComputerName, errc = os.Hostname()
	Key                = "ab9312a52781f4b7c7edf4341ef940daff94c567ffa503c3db8125fec68c4225" //encode key in bytes to string and keep as secret, put in a vault
	ActivationFile     = BASE_DIR + "/@ain8a.book"
	BlocklistFile      = BASE_DIR + "/@bft64.book"
	HostsFile          = BASE_DIR + "/@h5yg5.book"
)

// error constants
const (
	ERRORS_NULL                   = iota
	CHECK_ACTIVATION_READFAILURE  = iota
	CHECK_ACTIVATION_NOTACTIVATED = iota
)

// for scheduled tasks
// in minutes
const (
	ISKED_UPDATES       = 1 * 30
	ISKED_SYSTEM_STATUS = 4 * 60
	ISKED_SEND_HOSTS    = 2 * 60
)

// proxy commands
var (
	SetProxy   = "winhttp set proxy proxy-server=\"" + net.JoinHostPort(PROXY_HOST, PROXY_PORT) + "\" bypass-list=\"localhost\""
	ResetProxy = "winhttp reset proxy"
)

// global body structure map
type BodyStructure map[string]string
