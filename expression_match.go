package expression

import "regexp"

// RegularMatchString reports whether the string s
// contains any match of the regular expression pattern.
// More complicated queries need to use Compile and the full Regexp interface.
func RegularMatchString(patten string, s string) (bool, error) {
	return regexp.MatchString(patten, s)
}
