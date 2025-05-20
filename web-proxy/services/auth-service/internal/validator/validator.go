package validator

import "strings"

type Strategy string

const (
	StrategyAPIKey Strategy = "api_key"
	StrategyToken  Strategy = "token"
)

func Validate(strategy Strategy, input, secret string) bool {
	switch strategy {
	case StrategyAPIKey:
		return input == secret
	case StrategyToken:
		input = strings.TrimPrefix(input, "Bearer ")
		return input == secret
	default:
		return false
	}
}
