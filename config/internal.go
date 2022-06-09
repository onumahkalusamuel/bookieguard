package config

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/denisbrodbeck/machineid"
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
	MockData           = BASE_DIR + "/mockdata.book"
	Hosts              = []string{}
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
	ISKED_UPDATES      = 3 * 60
	ISKED_START_SERVER = 30
	ISKED_SEND_HOSTS   = 3 * 60
)

// server things
const (
	PROXY_HOST = "localhost"
	WEB_HOST   = "localhost"
	PROXY_PORT = "8088"
	WEB_PORT   = "8777"
)

// global body structure map
type BodyStructure map[string]string
