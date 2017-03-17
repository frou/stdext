package stdext

import (
	"errors"
	"regexp"
)

// ExtractNamedMatches tries to match the regular expression exp once in s and
// returns a map with keys which are the subexpression names in exp
// (?P<example>...) and values which are the matched values of those
// subexpressions.
func ExtractNamedMatches(
	exp string, s string) (map[string]string, error) {

	sl, err := extractSuccessiveNamedMatches(exp, s, 1)
	if err != nil {
		return nil, err
	}
	return sl[0], nil
}

// ExtractSuccessiveNamedMatches is like ExtractNamedMatches, except it will
// try to match the regular expression exp multiple times in s, and so returns
// a slice of maps rather than a single map.
func ExtractSuccessiveNamedMatches(
	exp string, s string) ([]map[string]string, error) {

	return extractSuccessiveNamedMatches(exp, s, -1)
}

func extractSuccessiveNamedMatches(
	exp string, s string, numSuccessive int) ([]map[string]string, error) {

	re, err := regexp.Compile(exp)
	if err != nil {
		return nil, err
	}

	subexpNames := re.SubexpNames()
	if len(subexpNames) < 1 {
		return nil, errors.New(
			"stdext: expected a regexp that uses named subexpressions")
	}

	var allResults []map[string]string

	for _, subexpMatches := range re.FindAllStringSubmatch(s, numSuccessive) {
		results := make(map[string]string)
		for i := 1; i < len(subexpNames); i++ {
			results[subexpNames[i]] = subexpMatches[i]
		}

		allResults = append(allResults, results)
	}

	if allResults == nil {
		return nil, errors.New("stdext: regexp did not match")
	}
	return allResults, nil
}
