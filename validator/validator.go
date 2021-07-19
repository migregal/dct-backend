package validator

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

const (
	DNSName  = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	WinPath  = `^([a-zA-Z]:\\)?(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	UnixPath = `^[.~]?(/[^/\x00]*)*/?$`
)

var (
	rxDNSName  = regexp.MustCompile(DNSName)
	rxWinPath  = regexp.MustCompile(WinPath)
	rxUnixPath = regexp.MustCompile(UnixPath)
)

const (
	// Unknown is unresolved OS type
	Unknown = iota
	// Win is Windows type
	Win
	// Unix is *nix OS types
	Unix
)

func IsIP(str string) bool {
	return net.ParseIP(str) != nil
}

func IsIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}

func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

func IsPort(str string) bool {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

func IsDNSName(str string) bool {
	if str == "" || len(strings.Replace(str, ".", "", -1)) > 255 {
		// constraints already violated
		return false
	}
	return !IsIP(str) && rxDNSName.MatchString(str)
}

func IsHost(str string) bool {
	return IsIP(str) || IsDNSName(str)
}
