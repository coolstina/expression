package expression

import "regexp"

// RegularReplaceMatch Use a regular expression to replace original string to novel string.
// The underlying call is regexp.Compile(patten).ReplaceAllString() method.
func RegularReplaceMatch(patten string, original string, replace string) (string, error) {
	regular, err := regexp.Compile(patten)
	if err != nil {
		return original, err
	}
	return regular.ReplaceAllString(original, replace), nil
}
