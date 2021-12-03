package expression

import (
	"regexp"
)

var (
	ipv4 = regexp.MustCompile(`\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}`)
)

// IsIPv4 regular expression is used to check whether the IP address is IPV4.
func IsIPv4(str string) bool {
	return ipv4.MatchString(str)
}
