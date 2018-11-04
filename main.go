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

type DecisionModuleResult struct {
	Passed bool
}

func GenerateRule(rules []Rule) DecisionModuleResult {
	passed := false
	for _, rule := range rules {
		fmt.Println(rule)
	}
	result := DecisionModuleResult{Passed: passed}
	return result
}
