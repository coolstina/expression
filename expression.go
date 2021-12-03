package expression

import (
	"regexp"
)

var (
	ipv4      = regexp.MustCompile(`^\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}$`)
	cellphone = regexp.MustCompile(`^(\+86(-|\s)?)?1[3-9][0-9]{9}$`)
	telephone = regexp.MustCompile(`^(\d{3,4}-)?[1-9][0-9]{4,8}$`)
)

// IsIPv4 regular expression is used to check whether the IP address is IPV4.
func IsIPv4(str string) bool {
	return ipv4.MatchString(str)
}

// IsCellphoneOfChina regular expression is used to check whether the telephone number is china.
func IsCellphoneOfChina(str string) bool {
	return cellphone.MatchString(str)
}

// IsTelephoneOfChina regular expression is used to check whether the telephone number is china.
func IsTelephoneOfChina(str string) bool {
	return telephone.MatchString(str)
}
