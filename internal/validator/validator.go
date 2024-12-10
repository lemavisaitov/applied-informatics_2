package validator

type Validator interface {
	Validate(input string) bool
	FindMatches(input string) []string
}
