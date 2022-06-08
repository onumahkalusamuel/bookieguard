package config

import (
	"net/http"
	"os"

	"github.com/denisbrodbeck/machineid"
)

var (
	AppKey             = "b00ki-GURd"
	HashedID, err      = machineid.ProtectedID(AppKey)
	ComputerName, errc = os.Hostname()
	Key                = "ab9312a52781f4b7c7edf4341ef940daff94c567ffa503c3db8125fec68c4225" //encode key in bytes to string and keep as secret, put in a vault
	ActivationFile     = "@ain8a.book"
	BlocklistFile      = "@bft64.book"
	Hosts              = []string{}
)

const (
	ERRORS_NULL                   = iota
	CHECK_ACTIVATION_READFAILURE  = iota
	CHECK_ACTIVATION_NOTACTIVATED = iota
)

// in minutes
const (
	ISKED_UPDATES      = 3 * 60
	ISKED_START_SERVER = 30
	ISKED_SEND_HOSTS   = 3 * 60
)

// server things
const (
	PROXY_HOST = "localhost" // leave empty for localhost
	WEB_HOST   = "localhost" // leave empty for localhost
	PROXY_PORT = "8088"
	WEB_PORT   = "8777"

	HTTP_BASE_DIR = http.Dir("./web/")
)

var PROXY_SERVER_HANDLE *http.Server

// var WEB_SERVER_HANDLE http.Server

type BodyStructure map[string]string
