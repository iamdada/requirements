package requirements

import "fmt"

type Rule struct {
	Name           string
	Variable       string
	Value          int32
	Comparison     string
	Interests      []string
	Decline_reason string
	Passed         bool
}

func GenerateRule(rules []Rule) Rule {
	rule := rules[0]

	fmt.Println(rule)
	return rule
}
