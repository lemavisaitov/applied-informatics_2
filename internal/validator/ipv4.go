package validator

import "regexp"

type IPv4Validator struct {
	regex *regexp.Regexp
}

func NewIPv4Validator() *IPv4Validator {
	const ipv4Regex = `\b((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\b`
	return &IPv4Validator{regex: regexp.MustCompile(ipv4Regex)}
}

func (v *IPv4Validator) Validate(input string) bool {
	return v.regex.MatchString(input)
}

func (v *IPv4Validator) FindMatches(input string) []string {
	return v.regex.FindAllString(input, -1)
}
